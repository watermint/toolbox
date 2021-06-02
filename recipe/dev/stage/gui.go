package stage

import (
	"encoding/base64"
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
	"github.com/watermint/toolbox/infra/ui/app_template"
	"github.com/watermint/toolbox/infra/ui/app_template_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"net/http"
)

type Command struct {
	Command string `uri:"command" binding:"required"`
}

func (z Command) EncodeCommandUrl() string {
	return base64.URLEncoding.EncodeToString([]byte(z.Command))
}

func (z Command) DecodeCommandName() (name string, err error) {
	nameRaw, err := base64.URLEncoding.DecodeString(z.Command)
	if err != nil {
		return "", err
	}
	return string(nameRaw), nil
}

type Page struct {
	Name    string
	Layouts []string
}

func (z Page) Apply(htp app_template.Template) error {
	l := esl.Default().With(esl.String("name", z.Name), esl.Strings("layouts", z.Layouts))

	if err := htp.Define(z.Name, z.Layouts...); err != nil {
		l.Debug("Unable to prepare templates", esl.Error(err))
		return err
	}
	l.Debug("The layout defined")
	return nil
}

type Server struct {
	ctl               app_control.Control
	expectedUserAgent string
}

func (z *Server) noSession(g *gin.Context) {
	l := z.ctl.Log()
	l.Info("noSession")
	g.HTML(
		http.StatusOK,
		"error",
		gin.H{},
	)
}
func (z *Server) home(g *gin.Context) {
	l := z.ctl.Log()
	l.Info("home")

	if g.Request.UserAgent() != z.expectedUserAgent {
		g.Redirect(
			http.StatusFound,
			"/no_session",
		)
		return
	}

	menu := make([]map[string]string, 0)
	menu = append(menu, map[string]string{
		"Uri":         "/catalogue",
		"Title":       "Commands",
		"Description": "Show available commands",
	})

	g.HTML(
		http.StatusOK,
		"home",
		gin.H{
			"Menu": menu,
		},
	)
}

func (z *Server) catalogue(g *gin.Context) {
	l := z.ctl.Log()
	l.Info("catalogue")
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
		if s.IsSecret() {
			continue
		}
		cat = append(cat, map[string]string{
			"Title":       s.CliPath(),
			"Description": ui.Text(s.Title()),
			"Uri":         "/command/" + Command{Command: s.CliPath()}.EncodeCommandUrl(),
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

func (z *Server) command(g *gin.Context) {
	l := z.ctl.Log()
	l.Info("command")
	ui := z.ctl.UI()

	cmd := &Command{}
	if err := g.ShouldBindUri(cmd); err != nil {
		g.HTML(
			http.StatusOK,
			"error",
			gin.H{},
		)
		return
	}
	cmdCliPath, err := cmd.DecodeCommandName()
	if err != nil {
		g.HTML(
			http.StatusOK,
			"error",
			gin.H{},
		)
		return
	}
	spec := app_catalogue.Current().RecipeSpec(cmdCliPath)

	cmdValues := make([]map[string]interface{}, 0)
	for _, valName := range spec.ValueNames() {
		// valSpec := spec.Value(valName)
		cmdValues = append(cmdValues, map[string]interface{}{
			"Name":    valName,
			"Desc":    ui.Text(spec.ValueDesc(valName)),
			"Default": spec.ValueDefault(valName),
		})
	}

	g.HTML(
		http.StatusOK,
		"command",
		gin.H{
			"Command":      cmdCliPath,
			"CommandTitle": ui.Text(spec.Title()),
			"CommandDesc":  ui.TextOrEmpty(spec.Desc()),
			"Values":       cmdValues,
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

	sessionAgent := fmt.Sprintf("%s/%s (Session%s)", app.Name, app.BuildId, sc_random.MustGetSecureRandomString(8))

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
	g.GET("/no_session", backend.noSession)
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
