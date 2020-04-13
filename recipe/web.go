package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_open"
	"github.com/watermint/toolbox/infra/web/web_handler"
	"github.com/watermint/toolbox/infra/web/web_job"
	"github.com/watermint/toolbox/infra/web/web_user"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Web struct {
	Port mo_int.RangeInt
}

func (z *Web) Preset() {
	z.Port.SetRange(1024, 65535, app.DefaultWebPort)
}

func (z *Web) Test(c app_control.Control) error {
	return rc_exec.Exec(c, z, rc_recipe.NoCustomValues)
}

func (z *Web) Exec(c app_control.Control) error {
	l := c.Log()
	repo, err := web_user.SingleUserRepository(c.Workspace())
	if err != nil {
		return err
	}
	rur := repo.(web_user.RootUserRepository)
	ru := rur.RootUser()

	if c.Feature().IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	hfs := c.(app_control.ControlHttpFileSystem)
	htp := hfs.Template()
	htr := htp.(render.HTMLRender)
	cl := c.(app_control_launcher.ControlLauncher)

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))
	//g.StaticFS("/assets", hfs.HttpFileSystem())
	g.HTMLRender = htr

	baseUrl := fmt.Sprintf("http://localhost:%d", z.Port)
	jobChan := make(chan *web_job.WebJobRun)

	wh := web_handler.NewHandler(
		c,
		htp,
		cl,
		baseUrl,
		jobChan,
	)
	wh.Setup(g)

	go web_job.Runner(c, jobChan)

	loginUrl := baseUrl + web_handler.WebPathLogin + "/" + ru.UserHash()

	l.Info("Login url", zap.String("url", loginUrl))

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		err = g.Run(fmt.Sprintf(":%d", z.Port))
		if err != nil {
			l.Error("Unable to start", zap.Error(err))
		}
	}()

	time.Sleep(2 * time.Second)
	l.Info("Trying open the browser")
	if !c.Feature().IsTest() {
		if err := ut_open.New().Open(loginUrl, true); err != nil {
			l.Warn("Unable to open the browser", zap.Error(err))
			l.Info("Please open this url on the browser", zap.String("url", loginUrl))
		}
		wg.Wait()
	}

	return nil
}
