package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/skratchdot/open-golang/open"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/web/web_handler"
	"github.com/watermint/toolbox/infra/web/web_job"
	"github.com/watermint/toolbox/infra/web/web_user"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"go.uber.org/zap"
	"sync"
	"time"
)

const (
	webPort = 7800
)

type Web struct {
	Port int
}

func (z *Web) Preset() {
	z.Port = webPort
}

func (z *Web) Test(c app_control.Control) error {
	return qt_endtoend.HumanInteractionRequired()
}

func (z *Web) Console() {
}

func (z *Web) Exec(c app_control.Control) error {
	l := c.Log()
	repo, err := web_user.SingleUserRepository(c.Workspace())
	if err != nil {
		return err
	}
	rur := repo.(web_user.RootUserRepository)
	ru := rur.RootUser()

	if c.IsProduction() {
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

	wh := web_handler.NewHanlder(
		c,
		htp,
		cl,
		baseUrl,
		jobChan,
	)
	wh.Setup(g)

	go web_job.Runner(c, jobChan)

	loginUrl := baseUrl + web_handler.WebPathLogin + "/" + ru.UserHash()

	c.Log().Info("Login url", zap.String("url", loginUrl))

	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		defer wg.Done()
		err = g.Run(fmt.Sprintf(":%d", z.Port))
		if err != nil {
			c.Log().Error("Unable to start", zap.Error(err))
		}
	}()

	time.Sleep(2 * time.Second)
	c.Log().Info("Trying open browser")
	open.Start(loginUrl)
	wg.Wait()

	return nil
}
