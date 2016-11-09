package dupload

import (
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial"
	"github.com/dropbox/dropbox-sdk-go-unofficial/files"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	// Hard limit of API spec: 150MB.
	UPLOAD_CHUNK_THRESHOLD int64 = 128 * 1048576 // 128MB
	UPLOAD_CHUNK_SIZE      int64 = 128 * 1048576 // 128MB
)

type UploadContext struct {
	localBasePath string

	LocalPath          string
	LocalRecursive     bool
	LocalFollowSymlink bool
	DropboxBasePath    string
	DropboxToken       string
}

type UploadSkipped struct {
	File   *UploadContext
	Reason error
}

func Upload(uc *UploadContext) {
	queue := make(chan *UploadContext)
	skipped := make(chan *UploadSkipped)
	control := make(chan string)

	go uploadRoutine(queue, skipped, control)

	base, err := os.Lstat(uc.LocalPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped", uc.LocalPath, err)
		return
	}

	if base.IsDir() {
		uc.localBasePath = uc.LocalPath
	} else {
		uc.localBasePath = filepath.Dir(uc.LocalPath)
	}

	queuePath(uc, queue)

	// Enqueue stop command
	control <- "upload-stop"

	// Wait for finish
	select {
	case <-control:
		seelog.Info("Upload finished")
	}

	go skippedFileReport(skipped, control)
	control <- "skipped-stop"

	select {
	case <-control:
		seelog.Trace("Skipped report finished")
	}
}

func skippedFileReport(skipped chan *UploadSkipped, control chan string) {
	for {
		select {
		case s := <-skipped:
			seelog.Warn("Skipped file: [%s] by reason [%s]", s.File.LocalPath, s.Reason)

		case cmd := <-control:
			seelog.Trace("Skipped record finished")
			control <- cmd + "-ack"
		}
	}
}

func uploadRoutine(c chan *UploadContext, skipped chan *UploadSkipped, control chan string) {
	for {
		select {
		case uc := <-c:
			err := upload(uc)
			if err != nil {
				skipped <- &UploadSkipped{
					File:   uc,
					Reason: err,
				}
			}

		case cmd := <-control:
			seelog.Trace("Finished")
			control <- cmd + "-ack"
			return
		}
	}
}

func upload(uc *UploadContext) error {
	seelog.Trace("Uploading file: ", uc.LocalPath)
	info, err := os.Lstat(uc.LocalPath)
	if err != nil {
		seelog.Warnf("Unable to acquire information about path [%s] by error [%s]. Skipped.", uc.LocalPath, err)
		return err
	}

	relative, err := filepath.Rel(uc.localBasePath, uc.LocalPath)
	if err != nil {
		seelog.Warnf("Unable to compute relative path [%s] by error [%s]. Skipped.", uc.LocalPath, err)
		return err
	}
	dropboxPath := filepath.Join(uc.DropboxBasePath, relative)

	seelog.Tracef("Start uploading local[%s] dropbox[%s] relative [%s]", uc.LocalPath, dropboxPath, relative)
	defer seelog.Tracef("Finished uploading local[%s] dropbox[%s] relative [%s]", uc.localBasePath, dropboxPath, relative)
	if info.Size() < UPLOAD_CHUNK_THRESHOLD {
		return uploadSingle(uc, info, dropboxPath)
	} else {
		return uploadChunked(uc, info, dropboxPath)
	}
}

func uploadSingle(uc *UploadContext, info os.FileInfo, dropboxPath string) error {
	client := dropbox.Client(uc.DropboxToken, dropbox.Options{})
	f, err := os.Open(uc.LocalPath)
	if err != nil {
		seelog.Warn("Unable to open file. Skipped.", uc.LocalPath, err)
		return err
	}

	ci := files.NewCommitInfo(dropboxPath)
	ci.ClientModified = info.ModTime().UTC()
	ci.Mode = &files.WriteMode{
		Tag: "overwrite",
	}

	res, err := client.Upload(ci, f)
	if err != nil {
		seelog.Warn("Unable to upload file.", uc.LocalPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", uc.LocalPath, dropboxPath, res.Id)
	return nil
}

func uploadChunked(uc *UploadContext, info os.FileInfo, dropboxPath string) error {
	seelog.Tracef("Chunked upload: %s", uc.LocalPath)
	client := dropbox.Client(uc.DropboxToken, dropbox.Options{})
	f, err := os.Open(uc.LocalPath)
	if err != nil {
		seelog.Warnf("Unable to open file [%s] by error [%v]. Skipped.", uc.LocalPath, err)
		return err
	}

	seelog.Tracef("Chunked Upload: Start session: %s", uc.LocalPath)
	r := io.LimitReader(f, UPLOAD_CHUNK_SIZE)
	session, err := client.UploadSessionStart(files.NewUploadSessionStartArg(), r)
	if err != nil {
		seelog.Warnf("Unable to create upload file [%s] by error [%v]", uc.LocalPath, err)
		return err
	}

	seelog.Tracef("Chunked Upload: Session started file [%s] session [%s]", uc.LocalPath, session.SessionId)

	var writtenBytes, totalBytes int64

	writtenBytes = UPLOAD_CHUNK_SIZE
	totalBytes = info.Size()

	for (totalBytes - writtenBytes) > UPLOAD_CHUNK_SIZE {
		seelog.Tracef("Chunked Upload: Append file [%s], session [%s], written [%d]", uc.LocalPath, session.SessionId, writtenBytes)

		cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
		aa := files.NewUploadSessionAppendArg(cursor)

		r = io.LimitReader(f, UPLOAD_CHUNK_SIZE)
		err := client.UploadSessionAppendV2(aa, r)
		if err != nil {
			seelog.Warnf("Unable to upload file [%s] caused by error [%v]", uc.LocalPath, err)
			return err
		}
		seelog.Tracef("Chunked Upload: Append (done): path [%s] session [%s] written [%d]", uc.LocalPath, session.SessionId, writtenBytes)
		writtenBytes += UPLOAD_CHUNK_SIZE
	}

	seelog.Tracef("Chunked Upload: Finish path[%s] sessoin[%s] written[%d]", uc.LocalPath, session.SessionId, writtenBytes)
	cursor := files.NewUploadSessionCursor(session.SessionId, uint64(writtenBytes))
	ci := files.NewCommitInfo(dropboxPath)
	ci.Path = dropboxPath
	ci.ClientModified = info.ModTime().UTC()
	ci.Mode = &files.WriteMode{
		Tag: "overwrite",
	}
	fa := files.NewUploadSessionFinishArg(cursor, ci)
	res, err := client.UploadSessionFinish(fa, f)
	if err != nil {
		seelog.Warnf("Unable to finish upload: path[%s] caused by error [%s]", dropboxPath, err)
		return err
	}
	seelog.Infof("File uploaded [%s] -> [%s] (%s)", uc.LocalPath, dropboxPath, res.Id)
	return nil
}

func queuePath(uc *UploadContext, c chan *UploadContext) {
	info, err := os.Lstat(uc.LocalPath)
	if err != nil {
		seelog.Warn("Unable to acquire information about path. Skipped.", uc.LocalPath)
		return
	}
	if info.Mode()&os.ModeSymlink == os.ModeSymlink && !uc.LocalFollowSymlink {
		seelog.Infof("Skipped (symlink): %s", uc.LocalPath)
		return
	}

	if info.IsDir() {
		queueDir(uc, c)
	} else {
		seelog.Trace("Queue Path: ", uc.LocalPath)
		c <- uc
	}
}

func queueDir(uc *UploadContext, c chan *UploadContext) error {
	list, err := ioutil.ReadDir(uc.LocalPath)
	if err != nil {
		seelog.Warnf("Unable to load directory [%s]. Skipped", uc.LocalPath)
	}
	for _, f := range list {
		localPath := filepath.Join(uc.LocalPath, f.Name())

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
			LocalPath:          localPath,
			LocalRecursive:     uc.LocalRecursive,
			LocalFollowSymlink: uc.LocalFollowSymlink,
			DropboxBasePath:    uc.DropboxBasePath,
			DropboxToken:       uc.DropboxToken,
		}

		queuePath(child, c)
	}
	return nil
}
