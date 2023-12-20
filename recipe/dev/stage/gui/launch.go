package gui

import (
	"fmt"
	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_gin"
	"github.com/watermint/toolbox/essentials/log/wrapper/lgw_print"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
)

const (
	serverPort = 7801
)

type Launch struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkExperimental
}

func (z *Launch) Preset() {
}

func (z *Launch) Exec(c app_control.Control) error {
	l := c.Log()

	a, _ := astilectron.New(lgw_print.New(l), astilectron.Options{
		AppName: app_definitions.Name,
	})
	defer a.Close()

	a.On(astilectron.EventNameAppEventReady, func(e astilectron.Event) (deleteListener bool) {
		l.Debug("Ready", esl.Any("event", e))
		return false
	})

	hfs := app_resource.Bundle().Web().HttpFileSystem()
	htp := app_template_impl.NewDev(hfs, c)
	htr := htp.(render.HTMLRender)
	if !c.Feature().IsDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	sessionAgent := fmt.Sprintf("%s/%s (Session%s)", app_definitions.Name, app_definitions.BuildId, sc_random.MustGetSecureRandomString(8))

	backend := &Server{
		ctl:               c,
		expectedUserAgent: sessionAgent,
	}
	g := gin.New()
	g.Use(lgw_gin.GinWrapper(l))
	g.Use(lgw_gin.GinRecovery(l))
	g.StaticFS("/assets", hfs)
	g.GET("/home", backend.home)
	g.GET("/command/:command", backend.command)
	g.GET("/catalogue", backend.catalogue)
	g.HTMLRender = htr

	pages := []Page{
		{
			Name:    "home",
			Layouts: []string{"layout/layout.html", "pages/home.html"},
		},
		{
			Name:    "catalogue",
			Layouts: []string{"layout/layout.html", "pages/catalogue.html"},
		},
		{
			Name:    "command",
			Layouts: []string{"layout/layout.html", "pages/command.html"},
		},
		{
			Name:    "error",
			Layouts: []string{"layout/simple.html", "pages/error.html"},
		},
	}
	for _, page := range pages {
		if err := page.Apply(htp); err != nil {
			return err
		}
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: g,
	}

	go func() {
		sErr := server.ListenAndServe()
		l.Info("Server finished", esl.Error(sErr))
	}()

	if err := a.Start(); err != nil {
		return err
	}

	w, _ := a.NewWindow(fmt.Sprintf("http://localhost:%d/catalogue", serverPort), &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(600),
		Width:  astikit.IntPtr(600),
		Load: &astilectron.WindowLoadOptions{
			UserAgent: sessionAgent,
		},
	})

	if err := w.Create(); err != nil {
		return err
	}

	a.Wait()
	return nil
}

func (z *Launch) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
