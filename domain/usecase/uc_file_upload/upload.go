package uc_file_upload

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/domain/service/sv_file_folder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

const (
	statusReportInterval = 15 * time.Second
)

type Upload interface {
	Upload(localPath string, dropboxPath string) (summary *UploadSummary, err error)
	Estimate(localPath string, dropboxPath string) (summary *UploadSummary, err error)
}

type UploadMO struct {
	ProgressUpload  app_msg.Message
	ProgressSummary app_msg.Message
}

func New(ctx api_context.Context, specs *rp_spec_impl.Specs, k app_kitchen.Kitchen, mo *UploadMO, opt ...UploadOpt) Upload {
	opts := &UploadOpts{
		ChunkSizeKb: 150 * 1024,
	}
	for _, o := range opt {
		o(opts)
	}
	return &uploadImpl{
		ctx:   ctx,
		specs: specs,
		opts:  opts,
		mo:    mo,
		k:     k,
	}
}

type Comparator interface {
	Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error)
}

type SizeComparator struct {
	l *zap.Logger
}

func (z SizeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		if f.Size == localFile.Size() {
			l.Debug("Same file size", zap.Int64("size", localFile.Size()))
			return true, nil
		}
		l.Debug("Size diff found", zap.Int64("localFileSize", localFile.Size()), zap.Int64("dbxFileSize", f.Size))
		return true, nil
	}
	l.Debug("Fallback")
	return false, nil
}

type TimeComparator struct {
	l *zap.Logger
}

func (z TimeComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lt := api_util.RebaseTime(localFile.ModTime())
		dt, err := api_util.Parse(f.ClientModified)
		if err != nil {
			l.Debug("Unable to parse client modified", zap.Error(err))
			return false, err
		}
		if lt.Equal(dt) {
			l.Debug("Same modified time", zap.String("clientModified", dt.String()))
			return true, nil
		}
		l.Debug("Modified time diff found",
			zap.String("localModTime", lt.String()),
			zap.String("dbxModTime", dt.String()),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

type HashComparator struct {
	l *zap.Logger
}

func (z HashComparator) Compare(localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	l := z.l.With(zap.String("localPath", localPath), zap.String("dbxPath", dbxEntry.PathDisplay()))
	if f, ok := dbxEntry.File(); ok {
		lch, err := api_util.ContentHash(localPath)
		if err != nil {
			l.Debug("Unable to calc local file content hash", zap.Error(err))
			return false, err
		}
		if lch == f.ContentHash {
			l.Debug("Same content hash", zap.String("hash", f.ContentHash))
			return true, nil
		}

		l.Debug("Content hash diff found",
			zap.String("localFileHash", lch),
			zap.String("dbxFileHash", f.ContentHash),
		)
		return false, nil
	}

	l.Debug("Fallback")
	return false, nil
}

// Returns true if it determined as same file
func Compare(l *zap.Logger, localPath string, localFile os.FileInfo, dbxEntry mo_file.Entry) (bool, error) {
	sc := &SizeComparator{l: l}
	tc := &TimeComparator{l: l}
	hc := &HashComparator{l: l}

	eq, err := sc.Compare(localPath, localFile, dbxEntry)
	if err != nil || !eq {
		return eq, err
	}
	eq, err = tc.Compare(localPath, localFile, dbxEntry)
	if err != nil {
		return eq, err
	}
	// determine as true, if same size & time
	if eq {
		return true, nil
	}

	// otherwise, compare content hash
	eq, err = hc.Compare(localPath, localFile, dbxEntry)
	return eq, err
}

type UploadOpt func(o *UploadOpts) *UploadOpts
type UploadOpts struct {
	Overwrite    bool
	ChunkSizeKb  int
	CreateFolder bool

	// TODO: should be extended for multiple folder name patterns
	ExcludeFolderName string
}

func ExcludeFolderName(folderName string) UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.ExcludeFolderName = folderName
		return o
	}
}
func ChunkSizeKb(size int) UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.ChunkSizeKb = size
		return o
	}
}
func Overwrite() UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.Overwrite = true
		return o
	}
}
func CreateFolder() UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.CreateFolder = true
		return o
	}
}

type UploadRow struct {
	File string `json:"file"`
	Size int64  `json:"size"`
}

type UploadSummary struct {
	UploadStart    time.Time `json:"upload_start"`
	UploadEnd      time.Time `json:"upload_end"`
	NumBytes       int64     `json:"num_bytes"`
	NumFilesError  int64     `json:"num_files_error"`
	NumFilesUpload int64     `json:"num_files_upload"`
	NumFilesSkip   int64     `json:"num_files_skip"`
	NumApiCall     int64     `json:"num_api_call"`
}

const (
	reportUpload  = "upload"
	reportSkip    = "skip"
	reportSummary = "summary"
)

func UploadReports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(
			reportUpload,
			rp_model.TransactionHeader(&UploadRow{}, &mo_file.ConcreteEntry{}),
			rp_model.HiddenColumns(
				"result.id",
				"result.tag",
			)),
		rp_spec_impl.Spec(
			reportSkip,
			rp_model.TransactionHeader(&UploadRow{}, &mo_file.ConcreteEntry{}),
			rp_model.HiddenColumns(
				"result.id",
				"result.tag",
			)),
		rp_spec_impl.Spec(reportSummary, &UploadSummary{}),
	}
}

type UploadStatus struct {
	summary UploadSummary
}

func (z *UploadStatus) error() {
	atomic.AddInt64(&z.summary.NumFilesError, 1)
}

func (z *UploadStatus) skip() {
	atomic.AddInt64(&z.summary.NumFilesSkip, 1)
}

func (z *UploadStatus) upload(size int64, chunkSize int) {
	atomic.AddInt64(&z.summary.NumBytes, size)
	atomic.AddInt64(&z.summary.NumFilesUpload, 1)

	apiCalls := int(math.Ceil(float64(size) / float64(chunkSize)))
	// Zero size file also consume API
	if size == 0 {
		apiCalls = 1
	}
	atomic.AddInt64(&z.summary.NumApiCall, int64(apiCalls))
}

type UploadWorker struct {
	dropboxBasePath string
	localBasePath   string
	localFilePath   string
	dbxEntry        mo_file.Entry
	ctx             api_context.Context
	ctl             app_control.Control
	up              sv_file_content.Upload
	mo              *UploadMO
	estimateOnly    bool
	repUpload       rp_model.Report
	repSkip         rp_model.Report
	opts            *UploadOpts
	status          *UploadStatus
}

func (z *UploadWorker) Exec() (err error) {
	ui := z.ctl.UI()
	upRow := &UploadRow{File: z.localFilePath}
	l := z.ctl.Log().With(
		zap.String("dropboxBasePath", z.dropboxBasePath),
		zap.String("localBasePath", z.localBasePath),
		zap.String("localFilePath", z.localFilePath),
	)
	l.Debug("Prepare upload")

	rel, err := ut_filepath.Rel(z.localBasePath, filepath.Dir(z.localFilePath))
	if err != nil {
		l.Debug("unable to calculate rel path", zap.Error(err))
		z.repUpload.Failure(err, upRow)
		z.status.error()
		return err
	}
	dp := mo_path.NewPath(z.dropboxBasePath)
	switch {
	case rel == ".":
		l.Debug("upload to base path")
	case strings.HasPrefix(rel, ".."):
		l.Debug("invalid rel path", zap.String("rel", rel))
		z.repUpload.Failure(errors.New("invalid path"), &UploadRow{File: z.localFilePath})
		z.status.error()
		return errors.New("invalid rel path")
	default:
		dp = dp.ChildPath(filepath.ToSlash(rel))
	}

	info, err := os.Lstat(z.localFilePath)
	if err != nil {
		z.repUpload.Failure(err, upRow)
		z.status.error()
		return err
	}
	upRow.Size = info.Size()

	// Verify proceed
	if z.dbxEntry != nil {
		same, err := Compare(l, z.localFilePath, info, z.dbxEntry)
		if err != nil {
			z.repUpload.Failure(err, upRow)
			z.status.error()
			return err
		}
		if same {
			//ui.Info("usecase.uc_file_upload.progress.skip", app_msg.P{
			//	"File": z.localFilePath,
			//})

			z.repSkip.Skip(app_msg.M("usecase.uc_file_upload.skip.file_exists"), upRow)
			z.status.skip()
			return nil
		}
	}

	if z.estimateOnly {
		z.status.upload(info.Size(), z.opts.ChunkSizeKb)
		l.Debug("Skip upload (estimate only)")
		return nil
	}

	ui.InfoM(z.mo.ProgressUpload.With("File", z.localFilePath))

	var entry mo_file.Entry
	if z.opts.Overwrite {
		entry, err = z.up.Overwrite(dp, z.localFilePath)
		if err != nil {
			z.repUpload.Failure(err, upRow)
			z.status.error()
			return err
		}
	} else {
		entry, err = z.up.Add(dp, z.localFilePath)
		if err != nil {
			z.repUpload.Failure(err, upRow)
			z.status.error()
			return err
		}
	}
	z.repUpload.Success(upRow, entry.Concrete())
	z.status.upload(info.Size(), z.opts.ChunkSizeKb)
	return nil
}

type uploadImpl struct {
	ctx   api_context.Context
	specs *rp_spec_impl.Specs
	opts  *UploadOpts
	mo    *UploadMO
	k     app_kitchen.Kitchen
}

func (z *uploadImpl) exec(localPath string, dropboxPath string, estimate bool) (summary *UploadSummary, err error) {
	l := z.k.Log().With(zap.String("localPath", localPath), zap.String("dropboxPath", dropboxPath), zap.Bool("estimate", estimate))
	l.Debug("execute")
	repUpload, err := z.specs.Open(reportUpload)
	if err != nil {
		return nil, err
	}
	defer repUpload.Close()

	repSkip, err := z.specs.Open(reportSkip)
	if err != nil {
		return nil, err
	}
	defer repSkip.Close()

	repSummary, err := z.specs.Open(reportSummary)
	if err != nil {
		return nil, err
	}
	defer repSummary.Close()

	status := &UploadStatus{
		summary: UploadSummary{
			UploadStart: time.Now(),
		},
	}

	go func() {
		for {
			time.Sleep(statusReportInterval)

			dur := time.Now().Sub(status.summary.UploadStart) / time.Second
			if dur == 0 {
				continue
			}

			kps := status.summary.NumBytes / int64(dur) / 1024

			z.k.UI().InfoM(z.mo.ProgressSummary.
				With("Time", time.Now().Truncate(time.Second).Format("15:04:05")).
				With("NumFileUpload", status.summary.NumFilesUpload).
				With("NumFileSkip", status.summary.NumFilesSkip).
				With("NumFileError", status.summary.NumFilesError).
				With("NumBytes", status.summary.NumBytes/1_048_576).
				With("Kps", kps).
				With("NumApiCall", status.summary.NumApiCall))
		}
	}()

	l.Debug("upload", zap.Int("chunkSize", z.opts.ChunkSizeKb))
	up := sv_file_content.NewUpload(z.ctx, sv_file_content.ChunkSizeKb(z.opts.ChunkSizeKb))
	q := z.k.NewQueue()

	info, err := os.Lstat(localPath)
	if err != nil {
		l.Debug("Unable to fetch info", zap.Error(err))
		return nil, err
	}

	createFolder := func(path string) error {
		ll := l.With(zap.String("path", path))
		ll.Debug("Prepare create folder")
		rel, err := ut_filepath.Rel(localPath, path)
		if err != nil {
			l.Debug("unable to calculate rel path", zap.Error(err))
			repUpload.Failure(err, &UploadRow{File: path})
			status.error()
			return err
		}
		if rel == "." {
			ll.Debug("Skip")
			return nil
		}

		folderPath := mo_path.NewPath(dropboxPath).ChildPath(filepath.ToSlash(rel))
		ll = ll.With(zap.String("folderPath", folderPath.Path()), zap.String("rel", rel))
		ll.Debug("Create folder")

		entry, err := sv_file_folder.New(z.ctx).Create(folderPath)
		if err != nil {
			if api_util.ErrorSummaryPrefix(err, "path/conflict/folder") {
				ll.Debug("The folder already exist, ignore it", zap.Error(err))
				return nil
			} else {
				ll.Debug("Unable to create folder", zap.Error(err))
				repUpload.Failure(err, &UploadRow{File: path})
				return err
			}
		}
		repUpload.Success(&UploadRow{File: path}, entry.Concrete())

		return nil
	}

	var scanFolder func(path string) error
	scanFolder = func(path string) error {
		ll := l.With(zap.String("path", path))

		nameLower := strings.ToLower(filepath.Base(path))
		if strings.ToLower(z.opts.ExcludeFolderName) == nameLower {
			ll.Debug("Skip folder by rule")
			return nil
		}

		ll.Debug("Scanning folder")
		localEntries, err := ioutil.ReadDir(path)
		if err != nil {
			ll.Debug("Unable to read dir", zap.Error(err))
			return err
		}
		localPathRel, err := ut_filepath.Rel(localPath, path)
		if err != nil {
			ll.Debug("Unable to calc rel path", zap.Error(err))
			return err
		}

		dbxPath := mo_path.NewPath(dropboxPath)
		if localPathRel != "." {
			dbxPath = dbxPath.ChildPath(filepath.ToSlash(localPathRel))
		}

		dbxEntries, err := sv_file.NewFiles(z.ctx).List(dbxPath)
		if err != nil {
			if api_util.ErrorSummaryPrefix(err, "path/not_found") {
				ll.Debug("Dropbox entry not found", zap.String("dbxPath", dbxPath.Path()), zap.Error(err))
				dbxEntries = make([]mo_file.Entry, 0)
			} else {
				ll.Debug("Unable to read Dropbox entries", zap.String("dbxPath", dbxPath.Path()), zap.Error(err))
				return err
			}
		}
		dbxEntryByName := mo_file.MapByNameLower(dbxEntries)

		numEntriesProceed := 0
		var lastErr error
		for _, e := range localEntries {
			p := filepath.Join(path, e.Name())
			if api_util.IsFileNameIgnored(p) {
				ll.Debug("Ignore file", zap.String("p", p))
				var ps int64 = 0
				pi, err := os.Lstat(p)
				if err == nil {
					ps = pi.Size()
				}
				status.skip()
				repSkip.Skip(
					app_msg.M("usecase.uc_file_upload.skip.dont_sync"),
					UploadRow{
						File: p,
						Size: ps,
					})
				continue
			}
			numEntriesProceed++
			if e.IsDir() {
				lastErr = scanFolder(filepath.Join(path, e.Name()))
			} else {
				dbxEntry := dbxEntryByName[strings.ToLower(e.Name())]
				ll.Debug("Enqueue", zap.String("p", p))
				q.Enqueue(&UploadWorker{
					dropboxBasePath: dropboxPath,
					localBasePath:   localPath,
					localFilePath:   p,
					dbxEntry:        dbxEntry,
					ctx:             z.ctx,
					ctl:             z.k.Control(),
					up:              up,
					mo:              z.mo,
					estimateOnly:    estimate,
					repUpload:       repUpload,
					repSkip:         repSkip,
					opts:            z.opts,
					status:          status,
				})
			}
		}
		l.Debug("folder scan finished", zap.Int("numEntriesProceed", numEntriesProceed), zap.Error(lastErr))
		if numEntriesProceed == 0 && z.opts.CreateFolder {
			l.Debug("Create folder for empty folder")
			return createFolder(path)
		}
		return lastErr
	}

	var lastErr error
	if info.IsDir() {
		lastErr = scanFolder(localPath)
	} else {
		q.Enqueue(&UploadWorker{
			dropboxBasePath: dropboxPath,
			localBasePath:   localPath,
			localFilePath:   localPath,
			ctx:             z.ctx,
			ctl:             z.k.Control(),
			up:              up,
			mo:              z.mo,
			estimateOnly:    estimate,
			repUpload:       repUpload,
			repSkip:         repSkip,
			opts:            z.opts,
			status:          status,
		})
	}

	q.Wait()

	status.summary.UploadEnd = time.Now()
	repSummary.Row(&status.summary)
	return &status.summary, lastErr
}

func (z *uploadImpl) Upload(localPath string, dropboxPath string) (summary *UploadSummary, err error) {
	return z.exec(localPath, dropboxPath, false)
}

func (z *uploadImpl) Estimate(localPath string, dropboxPath string) (summary *UploadSummary, err error) {
	return z.exec(localPath, dropboxPath, true)
}
