package _import

import (
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_file"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_desktop"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/domain/service/sv_file_folder"
	"github.com/watermint/toolbox/domain/service/sv_file_relocation"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/usecase/uc_file_upload"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/recpie/app_worker"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/util/ut_filepath"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/quality/scenario/qs_file"
	"go.uber.org/zap"
	"golang.org/x/sync/semaphore"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type ViaAppVO struct {
	Peer              app_conn.ConnUserFile
	DestDropboxPath   string
	LocalPath         string
	PseudoDesktopPath string
}

const (
	viaAppQueueSize          = 1000
	viaAppWatchInterval      = 10 * time.Second
	viaAppReportLocalScanner = "local_scanner"
	viaAppReportDbxScanner   = "dbx_scanner"
	viaAppReportCopier       = "copier"
	viaAppReportMover        = "mover"
)

var (
	semLocalScanner = semaphore.NewWeighted(viaAppQueueSize)
	semDbxScanner   = semaphore.NewWeighted(viaAppQueueSize)
	semCopier       = semaphore.NewWeighted(viaAppQueueSize)
	semMover        = semaphore.NewWeighted(viaAppQueueSize)
)

type ViaAppAccount struct {
}

type ViaAppHashToRel struct {
	htr      map[string]string
	htrMutex sync.Mutex
}

func (z *ViaAppHashToRel) Set(hash, rel string) {
	l := app_root.Log()
	l.Debug("Set", zap.String("hash", hash), zap.String("rel", rel))
	z.htrMutex.Lock()
	z.htr[hash] = rel
	z.htrMutex.Unlock()
}

func (z *ViaAppHashToRel) Delete(hash string) {
	l := app_root.Log()
	l.Debug("Delete", zap.String("hash", hash))
	z.htrMutex.Lock()
	delete(z.htr, hash)
	z.htrMutex.Unlock()
}

func (z *ViaAppHashToRel) Get(hash string) (string, bool) {
	l := app_root.Log()
	z.htrMutex.Lock()
	defer z.htrMutex.Unlock()
	r, ok := z.htr[hash]
	l.Debug("Get", zap.String("hash", hash), zap.String("rel", r), zap.Bool("ok", ok))
	return r, ok
}

type ViaAppStates struct {
	DbxWorkPath     mo_path.Path
	DesktopPath     string
	DesktopWorkPath string
	Backlogs        int64
	backlogMutex    sync.Mutex
}

func (z *ViaAppStates) AddBacklog() {
	z.backlogMutex.Lock()
	defer z.backlogMutex.Unlock()
	atomic.AddInt64(&z.Backlogs, 1)
}
func (z *ViaAppStates) ReleaseBacklog() {
	z.backlogMutex.Lock()
	defer z.backlogMutex.Unlock()

	if z.Backlogs > 0 {
		atomic.AddInt64(&z.Backlogs, -1)
	}
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

type ViaAppMoverInput struct {
	DbxFileName string `json:"dbx_file_name"`
}

func (z *ViaAppLocalScannerWorker) Exec() error {
	l := z.ctx.Log().With(zap.String("curLocalPath", z.curLocalPath))
	defer semLocalScanner.Release(1)
	defer z.st.ReleaseBacklog()
	lsIn := &ViaAppLocalScannerInput{Path: z.curLocalPath}

	l.Debug("Scanning local path")

	entries, err := ioutil.ReadDir(z.curLocalPath)
	if err != nil {
		l.Debug("Unable to read dir", zap.Error(err))
		z.reps.repLocalScanner.Failure(err, lsIn)
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
			if api_util.IsFileNameIgnored(e.Name()) {
				l.Debug("Ignore file", zap.Any("file", e))
			} else {
				l.Debug("Enqueue file", zap.Any("file", e))
				files = append(files, e)
			}
		}
	}

	if len(files) < 1 {
		l.Debug("No file found in the folder")
		return nil
	}

	l.Debug("Enqueue to dbx scanner", zap.Int("numFiles", len(files)))
	z.reps.repLocalScanner.Success(lsIn, nil)
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
	dsIn := &ViaAppDbxScannerInput{Path: z.curLocalPath}

	rel, err := ut_filepath.Rel(z.vo.LocalPath, z.curLocalPath)
	if err != nil {
		l.Error("Invalid local path", zap.Error(err))
		z.reps.repDbxScanner.Failure(err, dsIn)
		return err
	}
	dbxPath := mo_path.NewPath(z.vo.DestDropboxPath)
	if rel != "." {
		dbxPath = dbxPath.ChildPath(rel)
	}

	l.Info("Scanning dbx path", zap.String("dbxPath", dbxPath.Path()))
	entries, err := sv_file.NewFiles(z.ctx).List(dbxPath)
	if err != nil {
		if api_util.ErrorSummaryPrefix(err, "path/not_found") {
			l.Debug("Path not found in dest dropbox path")
			entries = make([]mo_file.Entry, 0)
		} else {
			l.Error("Failed to scan dbx path", zap.Error(err))
			return err
		}
	}

	requireUpdate := make(map[string]bool)
	nameToLocal := make(map[string]os.FileInfo)

	for _, f := range z.files {
		ln := strings.ToLower(f.Name())
		requireUpdate[ln] = true
		nameToLocal[ln] = f
	}

	for _, entry := range entries {
		if f, e := entry.File(); e {
			en := strings.ToLower(f.Name())
			if lf, ok := nameToLocal[en]; !ok {
				same, _ := uc_file_upload.Compare(l, filepath.Join(z.curLocalPath, lf.Name()), lf, f)
				if same {
					requireUpdate[en] = false
				}
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
		z.reps.repDbxScanner.Success(dsIn, copyIn)
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

	if z.copyIn.DbxFilePath == "" {
		return errors.New("empty dbx file path")
	}

	workCopyName := fmt.Sprintf("%x", sha256.Sum256([]byte(z.copyIn.DbxFilePath)))
	workCopyPath := filepath.Join(z.st.DesktopWorkPath, workCopyName)

	l := z.ctx.Log().With(
		zap.Any("copyIn", z.copyIn),
		zap.String("workCopyName", workCopyName),
		zap.String("workCopyPath", workCopyPath),
	)

	l.Info("Copying from local to work")

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

	// Update hash
	l.Debug("Update hash mapping")
	z.htr.Set(workCopyName, z.copyIn.DbxFilePath)
	z.st.AddBacklog()

	l.Debug("Copy finished", zap.Int64("writtenBytes", writtenBytes))
	z.reps.repCopier.Success(z.copyIn, &ViaAppCopierTransaction{LocalWorkCopyPath: workCopyPath})
	return nil
}

func viaAppEnqueueMover(
	k app_kitchen.Kitchen,
	ctx api_context.Context,
	htr *ViaAppHashToRel,
	vo *ViaAppVO,
	qs *ViaAppQueues,
	reps *ViaAppReports,
	st *ViaAppStates,
	moveIn *ViaAppMoverInput,
) {
	l := ctx.Log()
	l.Debug("Trying enqueue to mover queue")

	if err := semMover.Acquire(context.Background(), 1); err != nil {
		l.Error("Unable to acquire semaphore", zap.Error(err))
		return
	}

	qs.qMover.Enqueue(&ViaAppMoverWorker{
		k:      k,
		ctx:    ctx,
		htr:    htr,
		vo:     vo,
		qs:     qs,
		reps:   reps,
		st:     st,
		moveIn: moveIn,
	})
}

type ViaAppMoverWorker struct {
	k      app_kitchen.Kitchen
	ctx    api_context.Context
	htr    *ViaAppHashToRel
	vo     *ViaAppVO
	qs     *ViaAppQueues
	reps   *ViaAppReports
	st     *ViaAppStates
	moveIn *ViaAppMoverInput
}

func (z *ViaAppMoverWorker) Exec() error {
	defer semMover.Release(1)

	dbxDestPath, exist := z.htr.Get(z.moveIn.DbxFileName)
	l := z.ctx.Log().With(
		zap.Any("moveIn", z.moveIn),
		zap.String("dbxDestPath", dbxDestPath),
	)
	if !exist {
		l.Warn("Mapping not found")
		err := errors.New("path mapping not found")
		z.reps.repMover.Failure(err, z.moveIn)
		return err
	}
	defer z.st.ReleaseBacklog()

	l.Info("Moving from work to dest")
	src := z.st.DbxWorkPath.ChildPath(z.moveIn.DbxFileName)
	dst := mo_path.NewPath(dbxDestPath)

	movedEntry, err := sv_file_relocation.New(z.ctx).Move(src, dst)
	if err != nil {
		l.Debug("Unable to move", zap.Error(err))
		z.reps.repMover.Failure(err, z.moveIn)
		return err
	}

	l.Debug("Removing hash from the map")
	z.htr.Delete(z.moveIn.DbxFileName)

	z.reps.repMover.Success(z.moveIn, movedEntry.Concrete())
	return nil
}

type ViaAppWatcher struct {
	k    app_kitchen.Kitchen
	ctx  api_context.Context
	htr  *ViaAppHashToRel
	vo   *ViaAppVO
	qs   *ViaAppQueues
	reps *ViaAppReports
	st   *ViaAppStates
	wg   sync.WaitGroup
}

func (z *ViaAppWatcher) Watch() {
	l := z.k.Log()
	for {
		time.Sleep(viaAppWatchInterval)
		l.Debug("Watching loop", zap.Int64("backlogs", z.st.Backlogs))

		entries, err := sv_file.NewFiles(z.ctx).List(z.st.DbxWorkPath)
		if err != nil {
			l.Debug("Unable to list work path", zap.Error(err))
			continue
		}

		for _, entry := range entries {
			l.Info("Enqueue entry", zap.Any("entry", entry))
			viaAppEnqueueMover(
				z.k,
				z.ctx,
				z.htr,
				z.vo,
				z.qs,
				z.reps,
				z.st,
				&ViaAppMoverInput{DbxFileName: entry.Name()},
			)
		}

		if z.st.Backlogs < 1 {
			l.Debug("No more backlogs")
			z.wg.Done()
			return
		}
	}
}

type ViaApp struct {
}

func (ViaApp) Hidden() {
}

func (ViaApp) Console() {
}

func (z *ViaApp) Requirement() app_vo.ValueObject {
	return &ViaAppVO{}
}

func (z *ViaApp) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*ViaAppVO)
	l := k.Log()

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}
	profile, err := sv_profile.NewProfile(ctx).Current()
	if err != nil {
		return err
	}

	workPathRel := "tbx-file-import-viaapp/" + time.Now().Format("2006-01-02T15-04-05")
	dbxWorkPath := mo_path.NewPath("/" + workPathRel)
	dbxWorkFolder, err := sv_file_folder.New(ctx).Create(dbxWorkPath)
	if err != nil {
		return err
	}
	l.Debug("Work path created", zap.Any("workFolder", dbxWorkFolder))

	desktopPath := ""
	if vo.PseudoDesktopPath != "" {
		desktopPath = vo.PseudoDesktopPath
	} else {
		personal, business, err := sv_desktop.New().Lookup()
		if err != nil {
			l.Debug("Unable to find desktop path")
			return err
		}
		switch {
		case personal != nil && profile.AccountType != "business":
			l.Debug("User personal account", zap.Any("personal", personal), zap.String("accountType", profile.AccountType))
			desktopPath = personal.Path
		case business != nil && profile.AccountType == "business":
			l.Debug("User business account", zap.Any("personal", personal), zap.String("accountType", profile.AccountType))
			desktopPath = business.Path
		default:
			l.Warn("Account type mismatch or the desktop client does not linked to appropriate account")
			return errors.New("invalid account type")
		}
	}

	dpInfo, err := os.Lstat(desktopPath)
	if err != nil {
		l.Debug("Unable to retrieve desktop path info", zap.Error(err))
		return err
	}

	if !dpInfo.IsDir() {
		l.Warn("Given Dropbox Desktop path is not a directory", zap.String("path", desktopPath))
		return errors.New("desktop path is not a directory")
	}

	repSpecs := rp_spec_impl.New(z, k.Control())
	reps := &ViaAppReports{}
	reps.repLocalScanner, err = repSpecs.Open(viaAppReportLocalScanner)
	if err != nil {
		return err
	}
	defer reps.repLocalScanner.Close()

	reps.repDbxScanner, err = repSpecs.Open(viaAppReportDbxScanner)
	if err != nil {
		return err
	}
	defer reps.repDbxScanner.Close()

	reps.repCopier, err = repSpecs.Open(viaAppReportCopier)
	if err != nil {
		return err
	}
	defer reps.repCopier.Close()

	reps.repMover, err = repSpecs.Open(viaAppReportMover)
	if err != nil {
		return err
	}
	defer reps.repMover.Close()

	desktopWorkPath := filepath.Join(desktopPath, workPathRel)
	if err := os.MkdirAll(desktopWorkPath, 0755); err != nil {
		l.Debug("Unable to create folder", zap.Error(err))
		return err
	}

	htr := &ViaAppHashToRel{
		htr: make(map[string]string),
	}

	qs := &ViaAppQueues{
		qLocalScanner: k.NewQueue(),
		qDbxScanner:   k.NewQueue(),
		qCopier:       k.NewQueue(),
		qMover:        k.NewQueue(),
	}

	st := &ViaAppStates{
		DbxWorkPath:     dbxWorkPath,
		DesktopPath:     desktopPath,
		DesktopWorkPath: desktopWorkPath,
		Backlogs:        0,
	}

	watcher := ViaAppWatcher{
		k:    k,
		ctx:  ctx,
		htr:  htr,
		vo:   vo,
		qs:   qs,
		reps: reps,
		st:   st,
	}

	pseudoDesktop := &ViaAppPseudoDesktop{
		desktopPath:     desktopPath,
		desktopWorkPath: desktopWorkPath,
		dbxWorkPath:     dbxWorkPath,
		ctx:             ctx,
		specs:           repSpecs,
		k:               k,
		st:              st,
	}

	if vo.PseudoDesktopPath != "" {
		l.Debug("Run pseudo desktop")
		go pseudoDesktop.Run()
	}

	l.Debug("Enqueue local scanner")
	viaAppEnqueueLocalScanner(
		k,
		ctx,
		htr,
		vo,
		qs,
		reps,
		st,
		vo.LocalPath,
	)

	l.Debug("Launch watcher")
	watcher.wg.Add(1)
	go watcher.Watch()
	watcher.wg.Wait()

	l.Debug("Waiting for qLocalScanner")
	qs.qLocalScanner.Wait()
	l.Debug("Waiting for qDbxScanner")
	qs.qDbxScanner.Wait()
	l.Debug("Waiting for qCopier")
	qs.qCopier.Wait()
	l.Debug("Waiting for qMover")
	qs.qMover.Wait()

	return nil
}

func (z *ViaApp) Test(c app_control.Control) error {
	l := c.Log()
	vo := &ViaAppVO{}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		return qt_recipe.NotEnoughResource()
	}

	scenario := qs_file.Scenario{}
	if err := scenario.Create(); err != nil {
		return err
	}
	vo.LocalPath = scenario.LocalPath

	pseudoDesktop, err := ioutil.TempDir("", "pseudo-desktop")
	if err != nil {
		l.Error("unable to create temp dir", zap.Error(err))
		return err
	}
	vo.PseudoDesktopPath = pseudoDesktop
	vo.DestDropboxPath = "/" + qt_recipe.TestTeamFolderName + "/" + time.Now().Format("2006-01-02T15-04-05")

	if err = z.Exec(app_kitchen.NewKitchen(c, vo)); err != nil {
		return err
	}

	return nil
}

func (z *ViaApp) Reports() []rp_spec.ReportSpec {
	reps := make([]rp_spec.ReportSpec, 0)
	reps = append(reps,
		rp_spec_impl.Spec(
			viaAppReportLocalScanner,
			rp_model.TransactionHeader(&ViaAppLocalScannerInput{}, nil),
		),
		rp_spec_impl.Spec(
			viaAppReportDbxScanner,
			rp_model.TransactionHeader(&ViaAppDbxScannerInput{}, &ViaAppCopierInput{}),
		),
		rp_spec_impl.Spec(
			viaAppReportCopier,
			rp_model.TransactionHeader(&ViaAppCopierInput{}, &ViaAppCopierTransaction{}),
		),
		rp_spec_impl.Spec(
			viaAppReportMover,
			rp_model.TransactionHeader(&ViaAppMoverInput{}, &mo_file.ConcreteEntry{}),
		),
	)
	reps = append(reps, uc_file_upload.UploadReports()...)
	return reps
}

type ViaAppPseudoDesktop struct {
	desktopPath     string
	desktopWorkPath string
	dbxWorkPath     mo_path.Path
	ctx             api_context.Context
	specs           *rp_spec_impl.Specs
	st              *ViaAppStates
	k               app_kitchen.Kitchen
}

func (z *ViaAppPseudoDesktop) Run() {
	l := z.k.Log()
	up := uc_file_upload.New(z.ctx, z.specs, z.k)

	dbxToLocalSync := func() {
		dbxEntries, err := sv_file.NewFiles(z.ctx).List(z.dbxWorkPath)
		if err != nil {
			l.Debug("Unable to list folder", zap.Error(err))
			z.k.Control().Abort(app_control.Reason(app_control.FatalPanic))
		}

		localEntries, err := ioutil.ReadDir(z.desktopWorkPath)
		if err != nil {
			l.Debug("Unable to list folder", zap.Error(err))
			z.k.Control().Abort(app_control.Reason(app_control.FatalPanic))
		}

		for _, de := range dbxEntries {
			dn := strings.ToLower(de.Name())
			for _, le := range localEntries {
				ln := strings.ToLower(le.Name())
				if dn == ln {
					err := os.Remove(filepath.Join(z.desktopWorkPath, le.Name()))
					if err != nil {
						l.Debug("Unable to remove", zap.Error(err), zap.Any("localEntry", le), zap.Any("dropboxEntry", de))
						z.k.Control().Abort(app_control.Reason(app_control.FatalPanic))
					}
					break
				}
			}
		}
	}

	localToDbxSync := func() {
		localEntries, err := ioutil.ReadDir(z.desktopWorkPath)
		if err != nil {
			l.Debug("Unable to list folder", zap.Error(err))
			z.k.Control().Abort(app_control.Reason(app_control.FatalPanic))
		}

		if len(localEntries) < 1 {
			l.Debug("No entry found")
			return
		}
		summary, err := up.Upload(z.desktopWorkPath, z.dbxWorkPath.Path())
		if err != nil {
			l.Debug("Unable to upload")
			z.k.Control().Abort(app_control.Reason(app_control.FatalPanic))
		}
		l.Debug("Uploaded", zap.Any("summary", summary))
	}

	for {
		time.Sleep(10 * time.Second)
		dbxToLocalSync()
		localToDbxSync()
		if z.st.Backlogs < 1 {
			l.Debug("Finished")
			return
		}
	}
}
