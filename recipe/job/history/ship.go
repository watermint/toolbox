package history

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"os"
)

type Ship struct {
	Peer               rc_conn.ConnUserFile
	DropboxPath        mo_path.DropboxPath
	ProgressUploading  app_msg.Message
	ErrorFailedArchive app_msg.Message
	ErrorFailedUpload  app_msg.Message
	OperationLog       rp_model.TransactionReport
}

type ShipInfo struct {
	JobId      string `json:"job_id"`
	RecipeName string `json:"recipe_name"`
}

func (z *Ship) Exec(c app_control.Control) error {
	historian := app_job_impl.NewHistorian(c)
	histories := historian.Histories()
	l := c.Log()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	for _, h := range histories {
		if h.JobId() == c.Workspace().JobId() {
			l.Debug("Skip current job")
			continue
		}
		si := &ShipInfo{
			JobId:      h.JobId(),
			RecipeName: h.RecipeName(),
		}
		c.UI().Info(z.ProgressUploading.With("JobId", h.JobId()))
		path, err := h.Archive()
		if err != nil {
			l.Debug("Unable to archive", zap.Error(err), zap.Any("history", h))
			c.UI().Error(z.ErrorFailedArchive.With("JobId", h.JobId()).With("Error", err.Error()))
			z.OperationLog.Failure(err, si)
			continue
		}
		entry, err := sv_file_content.NewUpload(z.Peer.Context()).Add(z.DropboxPath, path)
		if err != nil {
			l.Debug("Unable to upload", zap.Error(err), zap.Any("history", h))
			c.UI().Error(z.ErrorFailedUpload.With("JobId", h.JobId()).With("Error", err.Error()))
			z.OperationLog.Failure(err, si)
			continue
		}
		if err = os.Remove(path); err != nil {
			l.Debug("Unable to remove archive", zap.Error(err), zap.String("path", path))
		}
		z.OperationLog.Success(si, entry.Concrete())
	}
	return nil
}

func (z *Ship) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Ship{}, func(r rc_recipe.Recipe) {
		m := r.(*Ship)
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("job-history-ship")
	})
	if e, _ := qt_recipe.RecipeError(c.Log(), err); e != nil {
		return err
	}

	return qt_errors.ErrorHumanInteractionRequired
}

func (z *Ship) Preset() {
	z.OperationLog.SetModel(&ShipInfo{}, &mo_file.ConcreteEntry{})
}
