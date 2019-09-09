package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/skratchdot/open-golang/open"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/web/web_handler"
	"github.com/watermint/toolbox/infra/web/web_job"
	"github.com/watermint/toolbox/infra/web/web_user"
	"go.uber.org/zap"
	"time"
)

const (
	webPort = 7800
)

type WebVO struct {
	Port int
}

type Web struct {
}

func (z *Web) Test(c app_control.Control) error {
	return nil
}

func (z *Web) Console() {
}

func (z *Web) Requirement() app_vo.ValueObject {
	return &WebVO{
		Port: webPort,
	}
}

func (z *Web) Exec(k app_kitchen.Kitchen) error {
	var vo interface{} = k.Value()
	wvo := vo.(*WebVO)

	l := k.Log()
	repo, err := web_user.SingleUserRepository(k.Control().Workspace())
	if err != nil {
		return err
	}
	rur := repo.(web_user.RootUserRepository)
	ru := rur.RootUser()

	if k.Control().IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	hfs := k.Control().(app_control.ControlHttpFileSystem)
	htp := hfs.Template()
	htr := htp.(render.HTMLRender)
	cl := k.Control().(app_control_launcher.ControlLauncher)

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))
	//g.StaticFS("/assets", hfs.HttpFileSystem())
	g.HTMLRender = htr

	baseUrl := fmt.Sprintf("http://localhost:%d", wvo.Port)
	jobChan := make(chan *web_job.WebJobRun)

	wh := web_handler.NewHanlder(
		k.Control(),
		htp,
		cl,
		baseUrl,
		jobChan,
	)
	wh.Setup(g)

	go web_job.Runner(k.Control(), jobChan)

	loginUrl := baseUrl + web_handler.WebPathLogin + "/" + ru.UserHash()

	k.Log().Info("Login url", zap.String("url", loginUrl))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))

	time.Sleep(2 * time.Second)

	open.Start(loginUrl)

	return nil
}
