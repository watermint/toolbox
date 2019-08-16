package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/domain/infra/api_auth"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/experimental/app_conn"
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
	webPathConnectStart    = "/connect/start"
	webPathConnectAuth     = "/connect/auth"
	webPathConnectFinish   = "/connect/finish"
	webPathForbidden       = "/error/forbidden"
	webPathServerError     = "/error/server"
	webPathAuthFailed      = "/error/auth_failed"
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

	baseUrl := fmt.Sprintf("http://localhost:%d", wvo.Port)

	wh := &WebHandler{
		control:     k.Control(),
		Template:    htp,
		Launcher:    cl,
		BaseUrl:     baseUrl,
		authForUser: make(map[string]api_auth.Web),
	}
	wh.Setup(g)

	loginUrl := baseUrl + webPathLogin + "/" + ru.UserHash()

	k.Log().Info("Login url", zap.String("url", loginUrl))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))

	return nil
}

type WebHandler struct {
	control     app_control.Control
	Template    app_template.Template
	Launcher    app_control_launcher.ControlLauncher
	BaseUrl     string
	Root        *app_recipe_group.Group
	authForUser map[string]api_auth.Web
}

func (z *WebHandler) auth(user app_user.User) api_auth.Web {
	if a, ok := z.authForUser[user.UserHash()]; ok {
		return a
	} else {
		a = api_auth_impl.NewWeb(z.control)
		z.authForUser[user.UserHash()] = a
		return a
	}
}

func (z *WebHandler) setupUrls(g *gin.Engine) {
	g.Use()

	g.GET(webPathLogin+"/:user_hash", z.Login)

	g.GET(webPathForbidden, z.Forbidden)
	g.GET(webPathServerError, z.ServerError)
	g.GET(webPathCommandNotFound, z.CommandNotFound)
	g.GET(webPathAuthFailed, z.AuthFailed)

	g.NoRoute(z.NotFound)
	z.Template.Define("error", "layout/layout.html", "pages/error.html")

	g.GET(webPathHome, z.Home)
	g.GET(webPathHome+"/:command", z.Home)
	g.GET(webPathHome+"/:command/:tokenType", z.Home)
	g.GET("param", z.renderRecipeParam)
	z.Template.Define("home-catalogue", "layout/layout.html", "pages/catalogue.html")
	z.Template.Define("home-recipe-conn", "layout/layout.html", "pages/recipe_conn.html")
	z.Template.Define("home-recipe-param", "layout/layout.html", "pages/recipe_param.html")

	g.GET(webPathRoot, z.Instruction)
	z.Template.Define(webPathRoot, "layout/layout.html", "pages/home.html")

	g.GET(webPathConnectStart, z.connectStart)
	g.GET(webPathConnectAuth, z.connectAuth)
}

func (z *WebHandler) setupCatalogue() {
	recipes := z.Launcher.Catalogue()
	z.Root = app_recipe_group.NewGroup([]string{}, "")
	for _, r := range recipes {
		_, ok := r.(app_recipe.SecretRecipe)
		if ok {
			continue
		}
		_, ok = r.(app_recipe.ConsoleRecipe)
		if ok {
			continue
		}

		z.Root.Add(r)
	}
}

func (z *WebHandler) findRecipe(cmd string) (grp *app_recipe_group.Group, rcp app_recipe.Recipe, err error) {
	cmdPath := strings.Split(cmd, "-")
	_, grp, rcp, _, err = z.Root.Select(cmdPath)

	if cmd == "" {
		grp = z.Root
		rcp = nil
		err = nil
	}
	return
}

func (z *WebHandler) recipeRequirements(rcp app_recipe.Recipe) (conns map[string]string, paramTypes map[string]string, paramDefaults map[string]interface{}) {
	conns = make(map[string]string)
	paramTypes = make(map[string]string)
	paramDefaults = make(map[string]interface{})

	var vo interface{} = rcp.Requirement()
	vc := app_vo_impl.NewValueContainer(vo)

	for k, v := range vc.Values {
		if d, ok := v.(bool); ok {
			paramTypes[k] = "bool"
			paramDefaults[k] = d
		} else if _, ok := v.(app_conn.ConnBusinessInfo); ok {
			conns[k] = api_auth.DropboxTokenBusinessInfo
		} else if _, ok := v.(app_conn.ConnBusinessFile); ok {
			conns[k] = api_auth.DropboxTokenBusinessFile
		} else if _, ok := v.(app_conn.ConnBusinessAudit); ok {
			conns[k] = api_auth.DropboxTokenBusinessAudit
		} else if _, ok := v.(app_conn.ConnBusinessMgmt); ok {
			conns[k] = api_auth.DropboxTokenBusinessManagement
		} else if _, ok := v.(app_conn.ConnUserFile); ok {
			conns[k] = api_auth.DropboxTokenFull
		}
	}
	return
}

func (z *WebHandler) Setup(g *gin.Engine) {
	z.setupCatalogue()
	z.setupUrls(g)
}

func (z *WebHandler) Login(g *gin.Context) {
	hash := g.Param("user_hash")
	repo, err := app_user.SingleUserRepository(z.control.Workspace())
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
			"Header": z.control.UI().Text("web.error.notfound.header"),
			"Detail": z.control.UI().Text("web.error.notfound.detail"),
		},
	)
}

func (z *WebHandler) ServerError(g *gin.Context) {
	g.HTML(
		http.StatusInternalServerError,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.server.header"),
			"Detail": z.control.UI().Text("web.error.server.detail"),
		},
	)
}
func (z *WebHandler) CommandNotFound(g *gin.Context) {
	g.HTML(
		http.StatusBadRequest,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.command_not_found.header"),
			"Detail": z.control.UI().Text("web.error.command_not_found.detail"),
		},
	)
}

func (z *WebHandler) AuthFailed(g *gin.Context) {
	g.HTML(
		http.StatusOK,
		"error",
		gin.H{
			"Header": z.control.UI().Text("web.error.auth_failed.header"),
			"Detail": z.control.UI().Text("web.error.auth_failed.detail"),
		},
	)
}

func (z *WebHandler) Forbidden(g *gin.Context) {
	g.HTML(
		http.StatusForbidden,
		webPathServerError,
		gin.H{
			"Header": z.control.UI().Text("web.error.forbidden.header"),
			"Detail": z.control.UI().Text("web.error.forbidden.detail"),
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

func (z *WebHandler) connectStart(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User) {
		//cmd := g.Query("command")
		tokenType := g.Query("tokenType")
		redirectUrl := z.BaseUrl + webPathConnectAuth
		_, url, err := z.auth(user).New(tokenType, redirectUrl)
		if err != nil {
			g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
			return
		}
		g.Redirect(http.StatusTemporaryRedirect, url)
	})
}

func (z *WebHandler) connectAuth(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User) {
		state := g.Query("state")
		code := g.Query("code")

		_, _, err := z.auth(user).Auth(state, code)
		if err != nil {
			g.Redirect(http.StatusTemporaryRedirect, webPathAuthFailed)
		} else {
			g.Redirect(http.StatusTemporaryRedirect, webPathConnectFinish)
		}
	})
}

func (z *WebHandler) Home(g *gin.Context) {
	z.withUser(g, func(g *gin.Context, user app_user.User) {
		cmd := g.Param("command")
		grp, rcp, err := z.findRecipe(cmd)

		switch {
		case err != nil:
			g.Redirect(http.StatusTemporaryRedirect, webPathCommandNotFound)

		case rcp != nil:
			// TODO: Breadcrumb list
			z.renderRecipeConn(g, cmd, rcp, user)

		case grp != nil:
			// TODO: Breadcrumb list
			z.renderCatalogue(g, cmd, grp)
		}
	})
}

func (z *WebHandler) renderRecipeConn(g *gin.Context, cmd string, rcp app_recipe.Recipe, user app_user.User) {
	l := z.control.Log().With(zap.String("cmd", cmd))
	reqConns, reqParams, _ := z.recipeRequirements(rcp)
	selectedConns := g.PostFormMap("Conn")
	orderedConnName := make([]string, 0)
	var nextConnName, nextConnType = "", ""

	for k := range reqConns {
		orderedConnName = append(orderedConnName, k)
	}
	sort.Strings(orderedConnName)

	for _, k := range orderedConnName {
		if v, ok := selectedConns[k]; !ok || v == "" {
			nextConnName = k
			nextConnType = reqConns[k]
		}
	}

	if nextConnName == "" { // no more required conns
		if len(reqParams) > 0 {
			// TODO forward to param rendering
			z.renderRecipeParam(g)
		} else {
			// TODO forward to confirm & run
		}
		return
	}

	existingConns, err := z.auth(user).List(nextConnType)
	if err != nil {
		l.Debug("Unable to list connections", zap.Error(err))
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
		return
	}
	listConns := make([]string, 0)
	connDesc := make(map[string]string)

	for _, e := range existingConns {
		listConns = append(listConns, e.PeerName)
		connDesc[e.PeerName] = e.Description
	}
	sort.Strings(listConns)

	g.HTML(
		http.StatusOK,
		"home-recipe-conn",
		gin.H{
			"Recipe":           cmd,
			"ExistingConns":    listConns,
			"ExistingConnDesc": connDesc,
			"SelectedConns":    selectedConns,
			"CurrentConn":      nextConnName,
			"CurrentConnType":  nextConnType,
		},
	)
}

func (z *WebHandler) renderRecipeParam(g *gin.Context) {
	g.HTML(
		http.StatusOK,
		"home-recipe-param",
		gin.H{
			"Recipe": "sharedfolder-list",
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
			"Description": z.control.UI().Text(grp.CommandDesc(g.Name).Key()),
			"Uri":         webPathHome + "/" + strings.Join(path, "-"),
		}
	}
	for name := range grp.Recipes {
		path := make([]string, 0)
		path = append(path, grp.Path...)
		path = append(path, name)

		dict[name] = gin.H{
			"Title":       name,
			"Description": z.control.UI().Text(grp.CommandDesc(name).Key()),
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
	repo, err := app_user.SingleUserRepository(z.control.Workspace())
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
