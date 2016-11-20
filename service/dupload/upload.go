package dupload

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial"
	"github.com/dropbox/dropbox-sdk-go-unofficial/files"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	// Hard limit of API spec: 150MB.
	UPLOAD_CHUNK_THRESHOLD int64 = 128 * 1048576 // 128MB
	UPLOAD_CHUNK_SIZE      int64 = 128 * 1048576 // 128MB
)

type UploadContext struct {
	localBasePath string
	localPath     string

	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
	DropboxToken       string
}

func Upload(srcPaths []string, baseUc *UploadContext, concurrency int) {
	queue := make(chan *UploadContext)
	wg := &sync.WaitGroup{}

	for i := 0; i < concurrency; i++ {
		go uploadRoutine(queue, wg)
	}
	for _, p := range srcPaths {
		scanPath(p, baseUc, wg, queue)
	}

	close(queue)
	wg.Wait()
}

func scanPath(srcPath string, baseUc *UploadContext, wg *sync.WaitGroup, queue chan *UploadContext) {
	uc := UploadContext{
		localPath:          srcPath,
		LocalRecursive:     baseUc.LocalRecursive,
		LocalFollowSymlink: baseUc.LocalFollowSymlink,
		DropboxBasePath:    baseUc.DropboxBasePath,
		DropboxToken:       baseUc.DropboxToken,
	}

	base, err := os.Lstat(uc.localPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped", uc.localPath, err)
		return
	}

	if base.IsDir() {
		uc.localBasePath = filepath.Clean(uc.localPath)
	} else {
		uc.localBasePath = filepath.Clean(filepath.Dir(uc.localPath))
	}

	seelog.Infof("Scanning files from: [%s]", uc.localBasePath)

	queuePath(&uc, queue, wg)
}

func uploadRoutine(uploadQueue chan *UploadContext, wg *sync.WaitGroup) {
	for uc := range uploadQueue {
		err := upload(uc)
		wg.Done()

		if err != nil {
			// TODO: err handling
		}
	}
}

func rebaseTime(t time.Time) time.Time {
	return t.Round(time.Second).UTC()
}

func upload(uc *UploadContext) error {
	seelog.Trace("Uploading file: ", uc.localPath)
	info, err := os.Lstat(uc.localPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped.", uc.localPath, err)
		return err
	}

	relative, err := filepath.Rel(uc.localBasePath, uc.localPath)
	if err != nil {
		seelog.Warnf("Unable to compute relative path [%s] by error [%s]. Skipped.", uc.localPath, err)
		return err
	}

	dropboxPath := filepath.ToSlash(filepath.Join(uc.DropboxBasePath, relative))

	seelog.Tracef("Start uploading local[%s] dropbox[%s] relative [%s]", uc.localPath, dropboxPath, relative)
	defer seelog.Tracef("Finished uploading local[%s] dropbox[%s] relative [%s]", uc.localBasePath, dropboxPath, relative)
	if info.Size() < UPLOAD_CHUNK_THRESHOLD {
		return uploadSingle(uc, info, dropboxPath)
	} else {
		return uploadChunked(uc, info, dropboxPath)
	}
}

func uploadSingle(uc *UploadContext, info os.FileInfo, dropboxPath string) error {
	client := dropbox.Client(uc.DropboxToken, dropbox.Options{})
	f, err := os.Open(uc.localPath)
	if err != nil {
		seelog.Warn("Unable to open file. Skipped.", uc.localPath, err)
		return err
	}

	ci := files.NewCommitInfo(dropboxPath)
	ci.ClientModified = rebaseTime(info.ModTime())
	ci.Mode = &files.WriteMode{
		Tag: "overwrite",
	}

	res, err := client.Upload(ci, f)
	if err != nil {
		seelog.Warn("Unable to upload file.", uc.localPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", uc.localPath, dropboxPath, res.Id)
	return nil
}

func uploadChunked(uc *UploadContext, info os.FileInfo, dropboxPath string) error {
	seelog.Tracef("Chunked upload: %s", uc.localPath)
	client := dropbox.Client(uc.DropboxToken, dropbox.Options{})
	f, err := os.Open(uc.localPath)
	if err != nil {
		seelog.Warnf("Unable to open file [%s] by error [%v]. Skipped.", uc.localPath, err)
		return err
	}

	seelog.Tracef("Chunked Upload: Start session: %s", uc.localPath)
	r := io.LimitReader(f, UPLOAD_CHUNK_SIZE)
	session, err := client.UploadSessionStart(files.NewUploadSessionStartArg(), r)
	if err != nil {
		seelog.Warnf("Unable to create upload file [%s] by error [%v]", uc.localPath, err)
		return err
	}

	seelog.Tracef("Chunked Upload: Session started file [%s] session [%s]", uc.localPath, session.SessionId)

	var writtenBytes, totalBytes int64

	writtenBytes = UPLOAD_CHUNK_SIZE
	totalBytes = info.Size()

	for (totalBytes - writtenBytes) > UPLOAD_CHUNK_SIZE {
		seelog.Tracef("Chunked Upload: Append file [%s], session [%s], written [%d]", uc.localPath, session.SessionId, writtenBytes)

		cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
		aa := files.NewUploadSessionAppendArg(cursor)

		r = io.LimitReader(f, UPLOAD_CHUNK_SIZE)
		err := client.UploadSessionAppendV2(aa, r)
		if err != nil {
			seelog.Warnf("Unable to upload file [%s] caused by error [%v]", uc.localPath, err)
			return err
		}
		seelog.Tracef("Chunked Upload: Append (done): path [%s] session [%s] written [%d]", uc.localPath, session.SessionId, writtenBytes)
		writtenBytes += UPLOAD_CHUNK_SIZE
	}

	seelog.Tracef("Chunked Upload: Finish path[%s] sessoin[%s] written[%d]", uc.localPath, session.SessionId, writtenBytes)
	cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
	ci := files.NewCommitInfo(dropboxPath)
	ci.Path = dropboxPath
	ci.ClientModified = rebaseTime(info.ModTime())
	ci.Mode = &files.WriteMode{
		Tag: "overwrite",
	}
	fa := files.NewUploadSessionFinishArg(cursor, ci)
	res, err := client.UploadSessionFinish(fa, f)
	if err != nil {
		seelog.Warnf("Unable to finish upload: path[%s] caused by error [%s]", dropboxPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", uc.localPath, dropboxPath, res.Id)
	return nil
}

func queuePath(uc *UploadContext, c chan *UploadContext, wg *sync.WaitGroup) {
	info, err := os.Lstat(uc.localPath)
	if err != nil {
		seelog.Warn("Unable to acquire information about path. Skipped.", uc.localPath)
		return
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink && !uc.LocalFollowSymlink {
		seelog.Infof("Skipped (symlink): %s", uc.localPath)
		return
	}

	if info.IsDir() {
		queueDir(uc, c, wg)
	} else {
		seelog.Trace("Queue Path: ", uc.localPath)
		c <- uc
		wg.Add(1)
	}
}

func queueDir(uc *UploadContext, c chan *UploadContext, wg *sync.WaitGroup) error {
	list, err := ioutil.ReadDir(uc.localPath)
	if err != nil {
		seelog.Warnf("Unable to load directory [%s]. Skipped", uc.localPath)
	}
	for _, f := range list {
		localPath := filepath.Join(uc.localPath, f.Name())

		if f.Mode()&os.ModeSymlink == os.ModeSymlink && !uc.LocalFollowSymlink {
			seelog.Infof("Skipped (symlink): %s", localPath)
			continue
		}
		if f.IsDir() && !uc.LocalRecursive {
			seelog.Infof("Skipped (without recursion): %s", localPath)
			continue
		}

		child := &UploadContext{
			localBasePath:      uc.localBasePath,
			localPath:          localPath,
			LocalRecursive:     uc.LocalRecursive,
			LocalFollowSymlink: uc.LocalFollowSymlink,
			DropboxBasePath:    uc.DropboxBasePath,
			DropboxToken:       uc.DropboxToken,
		}

		queuePath(child, c, wg)
	}
	return nil
}
