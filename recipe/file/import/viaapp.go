package _import

import (
	"context"
	"crypto/sha256"
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_worker"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
)

type ViaAppVO struct {
	Peer            app_conn.ConnUserFile
	DestDropboxPath string
	LocalPath       string
}

const (
	queueSize = 1000
)

var (
	semLocalScanner = semaphore.NewWeighted(queueSize)
	semDbxScanner   = semaphore.NewWeighted(queueSize)
	semCopier       = semaphore.NewWeighted(queueSize)
	semMover        = semaphore.NewWeighted(queueSize)
)

type ViaAppAccount struct {
}

type ViaAppHashToRel struct {
	htr      map[string]string
	htrMutex sync.Mutex
}

func (z *ViaAppHashToRel) Set(hash, rel string) {
	z.htrMutex.Lock()
	z.htr[hash] = rel
	z.htrMutex.Unlock()
}

func (z *ViaAppHashToRel) Delete(hash string) {
	z.htrMutex.Lock()
	delete(z.htr, hash)
	z.htrMutex.Unlock()
}

func (z *ViaAppHashToRel) Get(hash string) (string, bool) {
	z.htrMutex.Lock()
	defer z.htrMutex.Unlock()
	r, ok := z.htr[hash]
	return r, ok
}

type ViaAppStates struct {
	DbxWorkPath     mo_path.Path
	DesktopPath     string
	DesktopWorkPath string
	Backlogs        int64
}

func (z *ViaAppStates) AddBacklog() {
	atomic.AddInt64(&z.Backlogs, 1)
}
func (z *ViaAppStates) ReleaseBacklog() {
	atomic.AddInt64(&z.Backlogs, -1)
}

type ViaAppReports struct {
	repLocalScanner rp_model.Report
	repDbxScanner   rp_model.Report
	repCopier       rp_model.Report
	repMover        rp_model.Report
}

type ViaAppQueues struct {
	qLocalScanner app_worker.Queue
	qDbxScanner   app_worker.Queue
	qCopier       app_worker.Queue
	qMover        app_worker.Queue
}

type ViaAppLocalScannerWorker struct {
	k            app_kitchen.Kitchen
	ctx          api_context.Context
	htr          *ViaAppHashToRel
	vo           *ViaAppVO
	qs           *ViaAppQueues
	reps         *ViaAppReports
	st           *ViaAppStates
	curLocalPath string
}

func viaAppEnqueueLocalScanner(k app_kitchen.Kitchen,
	ctx api_context.Context,
	htr *ViaAppHashToRel,
	vo *ViaAppVO,
	qs *ViaAppQueues,
	reps *ViaAppReports,
	st *ViaAppStates,
	curLocalPath string) {

	l := ctx.Log()
	l.Debug("Trying enqueue to local scanner queue")

	if err := semLocalScanner.Acquire(context.Background(), 1); err != nil {
		l.Error("Unable to acquire semaphore", zap.Error(err))
		return
	}

	st.AddBacklog()
	qs.qLocalScanner.Enqueue(&ViaAppLocalScannerWorker{
		k:            k,
		ctx:          ctx,
		htr:          htr,
		vo:           vo,
		qs:           qs,
		reps:         reps,
		st:           st,
		curLocalPath: curLocalPath,
	})
}

type ViaAppLocalScannerInput struct {
	Path string `json:"path"`
}

type ViaAppDbxScannerInput struct {
	Path string `json:"path"`
}

type ViaAppCopierInput struct {
	LocalFilePath string `json:"local_file_path"`
	DbxFilePath   string `json:"dbx_file_path"`
}

type ViaAppCopierTransaction struct {
	LocalWorkCopyPath string `json:"local_work_copy_path"`
}

func (z *ViaAppLocalScannerWorker) Exec() error {
	l := z.ctx.Log().With(zap.String("curLocalPath", z.curLocalPath))
	defer semLocalScanner.Release(1)
	defer z.st.ReleaseBacklog()

	l.Debug("Scanning local path")

	entries, err := ioutil.ReadDir(z.curLocalPath)
	if err != nil {
		l.Debug("Unable to read dir", zap.Error(err))
		z.reps.repLocalScanner.Failure(err, &ViaAppLocalScannerInput{Path: z.curLocalPath})
		return err
	}

	files := make([]os.FileInfo, 0)

	for _, e := range entries {
		p := filepath.Join(z.curLocalPath, e.Name())
		if e.IsDir() {
			l.Debug("Enqueue local path", zap.String("path", p))
			viaAppEnqueueLocalScanner(
				z.k,
				z.ctx,
				z.htr,
				z.vo,
				z.qs,
				z.reps,
				z.st,
				p,
			)
		} else {
			l.Debug("Enqueue file", zap.Any("file", e))
			files = append(files, e)
		}
	}

	if len(files) < 1 {
		l.Debug("No file found in the folder")
		return nil
	}

	l.Debug("Enqueue to dbx scanner", zap.Int("numFiles", len(files)))
	viaAppEnqueueDbxScanner(
		z.k,
		z.ctx,
		z.htr,
		z.vo,
		z.qs,
		z.reps,
		z.st,
		z.curLocalPath,
		files,
	)
	return nil
}

func viaAppEnqueueDbxScanner(k app_kitchen.Kitchen,
	ctx api_context.Context,
	htr *ViaAppHashToRel,
	vo *ViaAppVO,
	qs *ViaAppQueues,
	reps *ViaAppReports,
	st *ViaAppStates,
	curLocalPath string,
	files []os.FileInfo) {

	l := ctx.Log()
	l.Debug("Trying enqueue to local scanner queue")

	if err := semDbxScanner.Acquire(context.Background(), 1); err != nil {
		l.Error("Unable to acquire semaphore", zap.Error(err))
		return
	}
	st.AddBacklog()
	qs.qDbxScanner.Enqueue(&ViaAppDbxScannerWorker{
		k:            k,
		ctx:          ctx,
		htr:          htr,
		vo:           vo,
		qs:           qs,
		reps:         reps,
		st:           st,
		curLocalPath: curLocalPath,
		files:        files,
	})
}

type ViaAppDbxScannerWorker struct {
	k            app_kitchen.Kitchen
	ctx          api_context.Context
	htr          *ViaAppHashToRel
	vo           *ViaAppVO
	qs           *ViaAppQueues
	reps         *ViaAppReports
	st           *ViaAppStates
	curLocalPath string
	files        []os.FileInfo
}

func (z *ViaAppDbxScannerWorker) Exec() error {
	l := z.ctx.Log().With(zap.String("curLocalPath", z.curLocalPath))
	defer semDbxScanner.Release(1)
	defer z.st.ReleaseBacklog()

	rel, err := ut_filepath.Rel(z.vo.LocalPath, z.curLocalPath)
	if err != nil {
		l.Error("Invalid local path", zap.Error(err))
		z.reps.repDbxScanner.Failure(err, &ViaAppDbxScannerInput{Path: z.curLocalPath})
		return err
	}
	dbxPath := mo_path.NewPath(z.vo.DestDropboxPath)
	if rel != "." {
		dbxPath = dbxPath.ChildPath(rel)
	}

	l.Debug("Scanning dbx path", zap.String("dbxPath", dbxPath.Path()))
	entries, err := sv_file.NewFiles(z.ctx).List(dbxPath)
	if err != nil {
		l.Error("Failed to scan dbx path", zap.Error(err))
		z.reps.repDbxScanner.Failure(err, &ViaAppDbxScannerInput{Path: z.curLocalPath})
		return err
	}

	requireUpdate := make(map[string]bool)
	nameToLocal := make(map[string]os.FileInfo)

	for _, f := range z.files {
		ln := strings.ToLower(f.Name())
		requireUpdate[ln] = true
		nameToLocal[ln] = f
	}

	compareFile := func(name string, lf os.FileInfo, df *mo_file.File) {
		lfp := filepath.Join(z.curLocalPath, lf.Name())
		ll := l.With(
			zap.String("localFilePath", lfp),
			zap.Int64("localFileSize", lf.Size()),
			zap.String("dbxFileContentHash", df.ContentHash),
			zap.Int64("dbxFileSize", df.Size),
		)

		if lf.Size() != df.Size {
			ll.Debug("Size difference found, leave mark as copy")
			requireUpdate[name] = true
			return
		}

		lt := lf.ModTime()
		dt, err := api_util.Parse(df.ClientModified)
		if err != nil {
			l.Debug("Unable to parse client modified time", zap.Error(err), zap.String("clientModified", df.ClientModified))
		} else {
			l.Debug("Compare time",
				zap.String("localFileTime", lt.String()),
				zap.String("dbxFileTime", dt.String()),
			)
			if lt.Equal(dt) {
				l.Debug("Skip copying (same mod time)")
				requireUpdate[name] = false
				return
			}
		}

		lch, err := api_util.ContentHash(lfp)
		if err != nil {
			l.Debug("Cannot compute hash, but leave mark as copy", zap.Error(err))
			requireUpdate[name] = true
			return

		} else {
			l.Debug("Computed hash of local file",
				zap.String("localFileContentHash", lch),
			)

			if lch == df.ContentHash {
				l.Debug("Skip copying (same content)")
				requireUpdate[name] = false
				return

			}

			l.Debug("Content diff found, leave mark as copy")
			requireUpdate[name] = true
		}
	}

	for _, entry := range entries {
		if f, e := entry.File(); e {
			en := strings.ToLower(f.Name())
			if lf, ok := nameToLocal[en]; !ok {
				compareFile(en, lf, f)
			}
		}
	}

	for name, update := range requireUpdate {
		lf := nameToLocal[name]

		if !update {
			continue
		}

		copyIn := &ViaAppCopierInput{
			LocalFilePath: filepath.Join(z.curLocalPath, lf.Name()),
			DbxFilePath:   dbxPath.ChildPath(lf.Name()).Path(),
		}

		l.Debug("Enqueue for copy", zap.String("localFile", lf.Name()), zap.Any("copyIn", copyIn))
		viaAppEnqueueCopier(
			z.k,
			z.ctx,
			z.htr,
			z.vo,
			z.qs,
			z.reps,
			z.st,
			copyIn,
		)
	}
	return nil
}

func viaAppEnqueueCopier(
	k app_kitchen.Kitchen,
	ctx api_context.Context,
	htr *ViaAppHashToRel,
	vo *ViaAppVO,
	qs *ViaAppQueues,
	reps *ViaAppReports,
	st *ViaAppStates,
	copyIn *ViaAppCopierInput,
) {
	l := ctx.Log()
	l.Debug("Trying enqueue to copier queue")

	if err := semCopier.Acquire(context.Background(), 1); err != nil {
		l.Error("Unable to acquire semaphore", zap.Error(err))
		return
	}

	st.AddBacklog()
	qs.qCopier.Enqueue(&ViaAppCopierWorker{
		k:      k,
		ctx:    ctx,
		htr:    htr,
		vo:     vo,
		qs:     qs,
		reps:   reps,
		st:     st,
		copyIn: copyIn,
	})
}

type ViaAppCopierWorker struct {
	k      app_kitchen.Kitchen
	ctx    api_context.Context
	htr    *ViaAppHashToRel
	vo     *ViaAppVO
	qs     *ViaAppQueues
	reps   *ViaAppReports
	st     *ViaAppStates
	copyIn *ViaAppCopierInput
}

func (z *ViaAppCopierWorker) Exec() error {
	defer semCopier.Release(1)
	defer z.st.ReleaseBacklog()

	workCopyName := fmt.Sprintf("%x", sha256.Sum256([]byte(z.copyIn.DbxFilePath)))
	workCopyPath := filepath.Join(z.st.DesktopWorkPath, workCopyName)

	l := z.ctx.Log().With(
		zap.Any("copyIn", z.copyIn),
		zap.String("workCopyName", workCopyName),
		zap.String("workCopyPath", workCopyPath),
	)
	l.Debug("Copying from local to work")

	l.Debug("Open source file")
	src, err := os.Open(z.copyIn.LocalFilePath)
	if err != nil {
		l.Debug("Unable to open local src file", zap.Error(err))
		z.reps.repCopier.Failure(err, z.copyIn)
		return err
	}
	defer src.Close()

	l.Debug("Create dest file")
	dst, err := os.Create(workCopyPath)
	if err != nil {
		l.Debug("unable to create dest file", zap.Error(err))
		z.reps.repCopier.Failure(err, z.copyIn)
		return err
	}

	l.Debug("Copy")
	writtenBytes, err := io.Copy(src, dst)
	if err != nil {
		l.Debug("Unable to copy", zap.Error(err))
		z.reps.repCopier.Failure(err, z.copyIn)
		return err
	}
	dst.Close()

	l.Debug("Copy finished", zap.Int64("writtenBytes", writtenBytes))
	z.reps.repCopier.Success(z.copyIn, &ViaAppCopierTransaction{LocalWorkCopyPath: workCopyPath})
	return nil
}
