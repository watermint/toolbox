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
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
)

type Server struct {
	ctl               app_control.Control
	expectedUserAgent string
}

func (z *Server) noSession(g *gin.Context) {
	esl.Default().Info("noSession")
	g.HTML(
		http.StatusOK,
		"error",
		gin.H{},
	)
}

func (z *Server) home(g *gin.Context) {
	esl.Default().Info("home")
	ui := z.ctl.UI()

	if g.Request.UserAgent() != z.expectedUserAgent {
		g.Redirect(
			http.StatusFound,
			"/no_session",
		)
		return
	}

	catRecipes := app_catalogue.Current().Recipes()
	cat := make([]map[string]string, 0)
	for _, r := range catRecipes {
		s := rc_spec.New(r)
		cat = append(cat, map[string]string{
			"Title":       s.CliPath(),
			"Description": ui.Text(s.Title()),
		})
	}

	g.HTML(
		http.StatusOK,
		"catalogue",
		gin.H{
			"Commands": cat,
		},
	)
}

type Gui struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkExperimental
}

func (z *Gui) Preset() {
}

func (z *Gui) Exec(c app_control.Control) error {
	l := c.Log()

	a, _ := astilectron.New(lgw_print.New(l), astilectron.Options{
		AppName: app.Name,
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

	sessionAgent := fmt.Sprintf("%s/%s (Session%s)", app.Name, app.Version, sc_random.MustGetSecureRandomString(8))

	backend := &Server{
		ctl:               c,
		expectedUserAgent: sessionAgent,
	}
	g := gin.New()
	g.Use(lgw_gin.GinWrapper(l))
	g.Use(lgw_gin.GinRecovery(l))
	g.StaticFS("/assets", hfs)
	g.GET("/catalogue", backend.home)
	g.GET("/no_session", backend.noSession)
	g.HTMLRender = htr
	if err := htp.Define("catalogue", "layout/simple.html", "pages/catalogue.html"); err != nil {
		l.Debug("Unable to prepare templates", esl.Error(err))
		return err
	}
	if err := htp.Define("error", "layout/simple.html", "pages/error.html"); err != nil {
		l.Debug("Unable to prepare templates", esl.Error(err))
		return err
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", 7800),
		Handler: g,
	}

	go func() {
		sErr := server.ListenAndServe()
		l.Info("Server finished", esl.Error(sErr))
	}()

	if err := a.Start(); err != nil {
		return err
	}

	w, _ := a.NewWindow("http://localhost:7800/catalogue", &astilectron.WindowOptions{
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

func (z *Gui) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
