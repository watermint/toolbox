package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_control_launcher"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_recipe_group"
	"github.com/watermint/toolbox/experimental/app_template"
	"github.com/watermint/toolbox/experimental/app_user"
	"github.com/watermint/toolbox/experimental/app_vo"
	"github.com/watermint/toolbox/experimental/app_vo_impl"
	"go.uber.org/zap"
	"net/http"
	"sort"
	"strings"
	"time"
)

const (
	webPort                = 7800
	webUserHashCookie      = "user_hash"
	webPathRoot            = "/"
	webPathLogin           = "/login"
	webPathHome            = "/home"
	webPathForbidden       = "/error/forbidden"
	webPathServerError     = "/error/server"
	webPathCommandNotFound = "/error/command_not_found"
)

type WebVO struct {
	Port int
}

func (z *WebVO) Validate(t app_vo.Validator) {
}

type Web struct {
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
	repo, err := app_user.SingleUserRepository(k.Control().Workspace())
	if err != nil {
		return err
	}
	rur := repo.(app_user.RootUserRepository)
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

	wh := &WebHandler{
		Kitchen:  k,
		Template: htp,
		Launcher: cl,
	}
	wh.Setup(g)

	loginUrl := fmt.Sprintf("http://localhost:%d%s/%s", wvo.Port, webPathLogin, ru.UserHash())

	k.Log().Info("Login url", zap.String("url", loginUrl))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))

	return nil
}

type WebHandler struct {
	Kitchen  app_kitchen.Kitchen
	Template app_template.Template
	Launcher app_control_launcher.ControlLauncher
}

func (z *WebHandler) Setup(g *gin.Engine) {
	g.GET(webPathLogin+"/:user_hash", z.Login)

	g.GET(webPathForbidden, z.Forbidden)
	g.GET(webPathServerError, z.ServerError)
	g.GET(webPathCommandNotFound, z.CommandNotFound)

	g.NoRoute(z.NotFound)
	z.Template.Define("error", "layout/layout.html", "pages/error.html")

	g.GET(webPathHome, z.Home)
	g.GET(webPathHome+"/:command", z.Home)
	z.Template.Define("home-catalogue", "layout/layout.html", "pages/catalogue.html")
	z.Template.Define("home-recipe", "layout/layout.html", "pages/recipe.html")

	g.GET(webPathRoot, z.Instruction)
	z.Template.Define(webPathRoot, "layout/layout.html", "pages/home.html")
}

func (z *WebHandler) Login(g *gin.Context) {
	hash := g.Param("user_hash")
	repo, err := app_user.SingleUserRepository(z.Kitchen.Control().Workspace())
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
	}
	_, err = repo.Resolve(hash)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
	}

	g.SetCookie(
		webUserHashCookie,
		hash,
		86400,
		"/",
		"localhost",
		false,
		true,
	)
	g.Redirect(http.StatusTemporaryRedirect, webPathHome)
}

func (z *WebHandler) NotFound(g *gin.Context) {
	g.HTML(
		http.StatusNotFound,
		"error",
		gin.H{
			"Header": z.Kitchen.UI().Text("web.error.notfound.header"),
			"Detail": z.Kitchen.UI().Text("web.error.notfound.detail"),
		},
	)
}

func (z *WebHandler) ServerError(g *gin.Context) {
	g.HTML(
		http.StatusInternalServerError,
		"error",
		gin.H{
			"Header": z.Kitchen.UI().Text("web.error.server.header"),
			"Detail": z.Kitchen.UI().Text("web.error.server.detail"),
		},
	)
}
func (z *WebHandler) CommandNotFound(g *gin.Context) {
	g.HTML(
		http.StatusBadRequest,
		"error",
		gin.H{
			"Header": z.Kitchen.UI().Text("web.error.command_not_found.header"),
			"Detail": z.Kitchen.UI().Text("web.error.command_not_found.detail"),
		},
	)
}

func (z *WebHandler) Forbidden(g *gin.Context) {
	g.HTML(
		http.StatusForbidden,
		webPathServerError,
		gin.H{
			"Header": z.Kitchen.UI().Text("web.error.forbidden.header"),
			"Detail": z.Kitchen.UI().Text("web.error.forbidden.detail"),
		},
	)
}

func (z *WebHandler) Instruction(g *gin.Context) {
	g.HTML(
		http.StatusOK,
		webPathRoot,
		gin.H{},
	)
	//g.Redirect(http.StatusTemporaryRedirect, "https://github.com/watermint/toolbox")
}

func (z *WebHandler) Home(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User) {
		recipes := z.Launcher.Catalogue()
		root := app_recipe_group.NewGroup([]string{}, "")
		for _, r := range recipes {
			_, ok := r.(app_recipe.SecretRecipe)
			if ok {
				continue
			}
			_, ok = r.(app_recipe.ConsoleRecipe)
			if ok {
				continue
			}

			root.Add(r)
		}

		cmd := g.Param("command")
		cmdPath := strings.Split(cmd, "-")
		_, grp, rcp, _, err := root.Select(cmdPath)

		if cmd == "" {
			grp = root
			rcp = nil
			err = nil
		}

		switch {
		case err != nil:
			g.Redirect(http.StatusTemporaryRedirect, webPathCommandNotFound)

		case rcp != nil:
			// TODO: Breadcrumb list
			z.renderRecipe(g, cmd, rcp)

		case grp != nil:
			// TODO: Breadcrumb list
			z.renderCatalogue(g, cmd, grp)
		}
	})
}

func (z *WebHandler) renderRecipe(g *gin.Context, cmd string, rcp app_recipe.Recipe) {
	var vo interface{} = rcp.Requirement()
	vc := app_vo_impl.NewValueContainer(vo)

	keys := make([]string, 0)
	for k := range vc.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	g.HTML(
		http.StatusOK,
		"home-recipe",
		gin.H{
			"Recipe": cmd,
			"Keys":   keys,
			"Values": vc.Values,
		},
	)
}

func (z *WebHandler) renderCatalogue(g *gin.Context, cmd string, grp *app_recipe_group.Group) {
	cmds := make([]string, 0)
	dict := make(map[string]gin.H)
	jobs := make([]gin.H, 0)

	for _, g := range grp.SubGroups {
		if g.IsSecret() {
			continue
		}

		path := make([]string, 0)
		path = append(path, grp.Path...)
		path = append(path, g.Name)

		dict[g.Name] = gin.H{
			"Title":       g.Name,
			"Description": z.Kitchen.UI().Text(grp.CommandDesc(g.Name).Key()),
			"Uri":         webPathHome + "/" + strings.Join(path, "-"),
		}
	}
	for name := range grp.Recipes {
		path := make([]string, 0)
		path = append(path, grp.Path...)
		path = append(path, name)

		dict[name] = gin.H{
			"Title":       name,
			"Description": z.Kitchen.UI().Text(grp.CommandDesc(name).Key()),
			"Uri":         webPathHome + "/" + strings.Join(path, "-"),
		}
	}

	for k := range dict {
		cmds = append(cmds, k)
	}
	sort.Strings(cmds)
	for _, c := range cmds {
		jobs = append(jobs, dict[c])
	}

	g.HTML(
		http.StatusOK,
		"home-catalogue",
		gin.H{
			"Detail": cmd,
			"Jobs":   jobs,
		},
	)
}

func (z *WebHandler) withUser(g *gin.Context, f func(g *gin.Context, user app_user.User)) {
	hash, err := g.Cookie(webUserHashCookie)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
		return
	}
	repo, err := app_user.SingleUserRepository(z.Kitchen.Control().Workspace())
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
		return
	}
	user, err := repo.Resolve(hash)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
		return
	}
	f(g, user)
}
