package stage

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_gin"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_print"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
)

type Gui struct {
}

func (z *Gui) Preset() {
}

func (z *Gui) home(g *gin.Context) {
	esl.Default().Info("home")
	g.HTML(
		http.StatusOK,
		"home",
		gin.H{
			"Detail": "GUI Proof of concept",
		},
	)
}

func (z *Gui) Exec(c app_control.Control) error {
	l := c.Log()

	a, _ := astilectron.New(lgw_print.New(l), astilectron.Options{
		AppName: app.Name,
	})
	defer a.Close()

	hfs := app_resource.Bundle().Web().HttpFileSystem()
	htp := app_template_impl.NewDev(hfs, c)
	htr := htp.(render.HTMLRender)
	if !c.Feature().IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}
	g := gin.New()
	g.Use(lgw_gin.GinWrapper(l))
	g.Use(lgw_gin.GinRecovery(l))
	g.StaticFS("/assets", hfs)
	g.GET("/", z.home)
	g.HTMLRender = htr
	if err := htp.Define("home", "layout/simple.html", "pages/home.html"); err != nil {
		l.Debug("Unable to prepare templates", esl.Error(err))
		return err
	}
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 7800),
		Handler: g,
	}
	go server.ListenAndServe()

	if err := a.Start(); err != nil {
		return err
	}

	w, _ := a.NewWindow("http://localhost:7800/", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
	})
	if err := w.Create(); err != nil {
		return err
	}

	a.Wait()
	return nil
}

func (z *Gui) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
