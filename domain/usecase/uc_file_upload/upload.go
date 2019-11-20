package uc_file_upload

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_content"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
)

type Upload interface {
	Upload(localPath string, dropboxPath string) (summary *UploadSummary, err error)
	Estimate(localPath string, dropboxPath string) (summary *UploadSummary, err error)
}

func New(ctx api_context.Context, specs *rp_spec_impl.Specs, k app_kitchen.Kitchen, opt ...UploadOpt) Upload {
	opts := &UploadOpts{
		ChunkSize: 150 * 1_048_576,
	}
	for _, o := range opt {
		o(opts)
	}
	return &uploadImpl{
		ctx:   ctx,
		specs: specs,
		opts:  opts,
		k:     k,
	}
}

type UploadOpt func(o *UploadOpts) *UploadOpts
type UploadOpts struct {
	Overwrite    bool
	ChunkSize    int
	CreateFolder bool
}

func ChunkSize(size int) UploadOpt {
	return func(o *UploadOpts) *UploadOpts {
		o.ChunkSize = size
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
	NumBytes       int64 `json:"num_bytes"`
	NumFilesError  int64 `json:"num_files_error"`
	NumFilesUpload int64 `json:"num_files_upload"`
	NumFilesSkip   int64 `json:"num_files_skip"`
	NumApiCall     int64 `json:"num_api_call"`
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
	ctx             api_context.Context
	ctl             app_control.Control
	up              sv_file_content.Upload
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

	rel, err := filepath.Rel(z.localBasePath, filepath.Dir(z.localFilePath))
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

	meta, err := sv_file.NewFiles(z.ctx).Resolve(sv_file_content.UploadPath(dp, info))
	if err != nil {
		l.Debug("Unable to resolve path", zap.Error(err))
	} else {
		f := meta.Concrete()
		if f.Size == info.Size() {
			hash, err := api_util.ContentHash(z.localFilePath)
			if err != nil {
				z.repUpload.Failure(err, upRow)
				z.status.error()
				return err
			}
			if f.ContentHash == hash {

				ui.Info("usecase.uc_file_upload.progress.skip", app_msg.P{
					"File": z.localFilePath,
				})

				z.repSkip.Skip(app_msg.M("usecase.uc_file_upload.skip.file_exists"), upRow)
				z.status.skip()
				return nil
			}
		}
	}

	if z.estimateOnly {
		z.status.upload(info.Size(), z.opts.ChunkSize)
		l.Debug("Skip upload (estimate only)")
		return nil
	}

	ui.Info("usecase.uc_file_upload.progress.upload", app_msg.P{
		"File": z.localFilePath,
	})

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
	z.status.upload(info.Size(), z.opts.ChunkSize)
	return nil
}

type uploadImpl struct {
	ctx   api_context.Context
	specs *rp_spec_impl.Specs
	opts  *UploadOpts
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

	status := &UploadStatus{}

	up := sv_file_content.NewUpload(z.ctx, sv_file_content.ChunkSize(int64(z.opts.ChunkSize)))
	q := z.k.NewQueue()

	info, err := os.Lstat(localPath)
	if err != nil {
		l.Debug("Unable to fetch info", zap.Error(err))
		return nil, err
	}

	createFolder := func(path string) error {
		ll := l.With(zap.String("path", path))
		ll.Debug("Create folder")
		// TODO: Implement for `sync up`

		return nil
	}

	var scanFolder func(path string) error
	scanFolder = func(path string) error {
		ll := l.With(zap.String("path", path))
		ll.Debug("Scanning folder")
		entries, err := ioutil.ReadDir(path)
		if err != nil {
			ll.Debug("Unable to read dir", zap.Error(err))
			return err
		}
		numEntriesProceed := 0
		var lastErr error
		for _, e := range entries {
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
				lastErr = scanFolder(e.Name())
			} else {
				ll.Debug("Enqueue", zap.String("p", p))
				q.Enqueue(&UploadWorker{
					dropboxBasePath: dropboxPath,
					localBasePath:   localPath,
					localFilePath:   p,
					ctx:             z.ctx,
					ctl:             z.k.Control(),
					up:              up,
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
			estimateOnly:    estimate,
			repUpload:       repUpload,
			repSkip:         repSkip,
			opts:            z.opts,
			status:          status,
		})
	}

	q.Wait()

	repSummary.Row(&status.summary)
	return &status.summary, lastErr
}

func (z *uploadImpl) Upload(localPath string, dropboxPath string) (summary *UploadSummary, err error) {
	return z.exec(localPath, dropboxPath, false)
}

func (z *uploadImpl) Estimate(localPath string, dropboxPath string) (summary *UploadSummary, err error) {
	return z.exec(localPath, dropboxPath, true)
}
