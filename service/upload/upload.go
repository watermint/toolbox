package upload

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/files"
	"github.com/watermint/bwlimit"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/integration/sdk"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	// Hard limit of API spec: 150MB.
	UPLOAD_CHUNK_THRESHOLD int64 = 128 * 1048576 // 128MB
	UPLOAD_CHUNK_SIZE      int64 = 128 * 1048576 // 128MB
)

type UploadContext struct {
	Infra              *infra.InfraOpts
	LocalPaths         []string
	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
	DropboxToken       string
	DeleteAfterUpload  bool
	BandwidthLimit     int
	Concurrency        int

	uploadQueue chan *scanContext
	throttle    *bwlimit.Bwlimit
	backlog     *sync.WaitGroup
}

type scanContext struct {
	localBasePath string
	localPath     string
	uploadContext *UploadContext
}

func (uc *UploadContext) Upload() {
	bw := bwlimit.NewBwlimit(uc.BandwidthLimit, true)
	uc.uploadQueue = make(chan *scanContext)
	uc.backlog = &sync.WaitGroup{}
	uc.throttle = &bw

	for i := 0; i < uc.Concurrency; i++ {
		go uc.uploadRoutine()
	}
	for _, p := range uc.LocalPaths {
		uc.scanPath(p)
	}

	close(uc.uploadQueue)

	uc.backlog.Wait()
	uc.throttle.Wait()
}

func (uc *UploadContext) scanPath(srcPath string) {
	sc := scanContext{
		localPath:     srcPath,
		uploadContext: uc,
	}

	base, err := os.Lstat(sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped", sc.localPath, err)
		return
	}

	if base.IsDir() {
		sc.localBasePath = filepath.Clean(sc.localPath)
	} else {
		sc.localBasePath = filepath.Clean(filepath.Dir(sc.localPath))
	}

	seelog.Infof("Scanning files from: [%s]", sc.localBasePath)

	sc.queuePath()
}

func (uc *UploadContext) uploadRoutine() {
	for sc := range uc.uploadQueue {
		err := sc.upload()
		uc.backlog.Done()

		if err != nil {
			// TODO: err handling
		}
	}
}

func (sc *scanContext) upload() error {
	seelog.Trace("Uploading file: ", sc.localPath)
	info, err := os.Lstat(sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped.", sc.localPath, err)
		return err
	}

	relative, err := filepath.Rel(sc.localBasePath, sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to compute relative path [%s] by error [%s]. Skipped.", sc.localPath, err)
		return err
	}

	dropboxPath := filepath.ToSlash(filepath.Join(sc.uploadContext.DropboxBasePath, relative))
	if !strings.HasPrefix(dropboxPath, "/") {
		// Dropbox path must be start from '/'
		dropboxPath = "/" + dropboxPath
	}

	seelog.Tracef("Start uploading local[%s] dropbox[%s] relative [%s]", sc.localPath, dropboxPath, relative)
	defer seelog.Tracef("Finished uploading local[%s] dropbox[%s] relative [%s]", sc.localBasePath, dropboxPath, relative)
	if info.Size() < UPLOAD_CHUNK_THRESHOLD {
		return sc.uploadSingle(info, dropboxPath)
	} else {
		return sc.uploadChunked(info, dropboxPath)
	}
}

func (sc *scanContext) uploadSingle(info os.FileInfo, dropboxPath string) error {
	config := dropbox.Config{Token: sc.uploadContext.DropboxToken, Verbose: false}
	client := files.New(config)
	f, err := os.Open(sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to open file. Skipped. localPath[%s] error[%s]", sc.localPath, err)
		return err
	}
	defer f.Close()
	bwf := sc.uploadContext.throttle.Reader(f)

	ci := files.NewCommitInfo(dropboxPath)
	ci.ClientModified = sdk.RebaseTimeForAPI(info.ModTime())

	res, err := client.Upload(ci, bwf)
	if err != nil {
		seelog.Warnf("Unable to upload file. path[%s] error[%s]", sc.localPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", sc.localPath, dropboxPath, res.Id)

	return nil
}

func (sc *scanContext) uploadChunked(info os.FileInfo, dropboxPath string) error {
	seelog.Tracef("Chunked upload: %s", sc.localPath)
	config := dropbox.Config{Token: sc.uploadContext.DropboxToken, Verbose: false}
	client := files.New(config)
	f, err := os.Open(sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to open file [%s] by error [%v]. Skipped.", sc.localPath, err)
		return err
	}
	defer f.Close()

	seelog.Tracef("Chunked Upload: Start session: %s", sc.localPath)
	r := sc.uploadContext.throttle.Reader(io.LimitReader(f, UPLOAD_CHUNK_SIZE))
	session, err := client.UploadSessionStart(files.NewUploadSessionStartArg(), r)
	if err != nil {
		seelog.Warnf("Unable to create upload file [%s] by error [%v]", sc.localPath, err)
		return err
	}

	seelog.Tracef("Chunked Upload: Session started file [%s] session [%s]", sc.localPath, session.SessionId)

	var writtenBytes, totalBytes int64

	writtenBytes = UPLOAD_CHUNK_SIZE
	totalBytes = info.Size()

	for (totalBytes - writtenBytes) > UPLOAD_CHUNK_SIZE {
		seelog.Tracef("Chunked Upload: Append file [%s], session [%s], written [%d]", sc.localPath, session.SessionId, writtenBytes)

		cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
		aa := files.NewUploadSessionAppendArg(cursor)

		r = sc.uploadContext.throttle.Reader(io.LimitReader(f, UPLOAD_CHUNK_SIZE))
		err := client.UploadSessionAppendV2(aa, r)
		if err != nil {
			seelog.Warnf("Unable to upload file [%s] caused by error [%v]", sc.localPath, err)
			return err
		}
		seelog.Tracef("Chunked Upload: Append (done): path [%s] session [%s] written [%d]", sc.localPath, session.SessionId, writtenBytes)
		writtenBytes += UPLOAD_CHUNK_SIZE
	}

	seelog.Tracef("Chunked Upload: Finish path[%s] sessoin[%s] written[%d]", sc.localPath, session.SessionId, writtenBytes)
	cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
	ci := files.NewCommitInfo(dropboxPath)
	ci.Path = dropboxPath
	ci.ClientModified = sdk.RebaseTimeForAPI(info.ModTime())
	fa := files.NewUploadSessionFinishArg(cursor, ci)
	res, err := client.UploadSessionFinish(fa, sc.uploadContext.throttle.Reader(f))
	if err != nil {
		seelog.Warnf("Unable to finish upload: path[%s] caused by error [%s]", dropboxPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", sc.localPath, dropboxPath, res.Id)
	return nil
}

func (sc *scanContext) queuePath() {
	info, err := os.Lstat(sc.localPath)
	if err != nil {
		seelog.Warn("Unable to acquire information about path. Skipped.", sc.localPath)
		return
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink && !sc.uploadContext.LocalFollowSymlink {
		seelog.Infof("Skipped (symlink): %s", sc.localPath)
		return
	}
	if info.IsDir() {
		sc.queueDir()
	} else {
		seelog.Trace("Queue Path: ", sc.localPath)
		sc.uploadContext.uploadQueue <- sc
		sc.uploadContext.backlog.Add(1)
	}
}

func (sc *scanContext) queueDir() error {
	list, err := ioutil.ReadDir(sc.localPath)
	if err != nil {
		seelog.Warnf("Unable to load directory [%s]. Skipped", sc.localPath)
	}
	for _, f := range list {
		localPath := filepath.Join(sc.localPath, f.Name())

		if f.Mode()&os.ModeSymlink == os.ModeSymlink && !sc.uploadContext.LocalFollowSymlink {
			seelog.Infof("Skipped (symlink): %s", localPath)
			continue
		}
		if f.IsDir() && !sc.uploadContext.LocalRecursive {
			seelog.Infof("Skipped (without recursion): %s", localPath)
			continue
		}

		child := &scanContext{
			uploadContext: sc.uploadContext,
			localBasePath: sc.localBasePath,
			localPath:     localPath,
		}

		child.queuePath()
	}
	return nil
}
