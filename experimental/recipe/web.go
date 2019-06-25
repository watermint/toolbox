package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/watermint/toolbox/experimental/app_control"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_template"
	"github.com/watermint/toolbox/experimental/app_user"
	"github.com/watermint/toolbox/experimental/app_vo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	webPort            = 7800
	webUserHashCookie  = "user_hash"
	webPathRoot        = "/"
	webPathLogin       = "/login"
	webPathHome        = "/home"
	webPathForbidden   = "/error/forbidden"
	webPathServerError = "/error/server"
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

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))
	//g.StaticFS("/assets", hfs.HttpFileSystem())
	g.HTMLRender = htr

	wh := &WebHandler{
		Kitchen:  k,
		Template: htp,
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
}

func (z *WebHandler) Setup(g *gin.Engine) {
	g.GET(webPathLogin+"/:user_hash", z.Login)

	g.GET(webPathForbidden, z.Forbidden)
	g.GET(webPathServerError, z.ServerError)
	g.NoRoute(z.NotFound)
	z.Template.Define("error", "layout/layout.html", "pages/error.html")

	g.GET(webPathHome, z.Home)
	z.Template.Define(webPathHome, "layout/layout.html", "pages/catalogue.html")

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
		g.HTML(
			http.StatusOK,
			webPathHome,
			gin.H{"Jobs": []gin.H{
				{
					"Title":       "member remove",
					"Description": "Remove members",
				},
				{
					"Title":       "member update",
					"Description": "Update members",
				},
				{
					"Title":       "member detach",
					"Description": "Detach members",
				},
				{
					"Title":       "member invite",
					"Description": "Invite members",
				},
			}},
		)
	})
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
