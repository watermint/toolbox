package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/experimental/app_msg"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_user"
	"github.com/watermint/toolbox/experimental/app_vo"
	"go.uber.org/zap"
	"net/http"
	"time"
)

const (
	webPort            = 7800
	webUserHashCookie  = "user_hash"
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

func (z *Web) Exec(k app_recipe.Kitchen) error {
	var vo interface{} = k.Value()
	wvo := vo.(*WebVO)

	l := k.Log()
	repo, err := app_user.SingleUserRepository(k.Control().Workspace())
	if err != nil {
		return err
	}
	rur := repo.(app_user.RootUserRepository)
	ru := rur.RootUser()

	if app.AppHash != "" {
		gin.SetMode(gin.ReleaseMode)
	}

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))

	wh := &WebHandler{
		Kitchen: k,
	}
	wh.Setup(g)

	loginUrl := fmt.Sprintf("http://localhost:%d%s/%s", wvo.Port, webPathLogin, ru.UserHash())

	k.Log().Info("Login url", zap.String("url", loginUrl))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))

	return nil
}

type WebHandler struct {
	Kitchen app_recipe.Kitchen
}

func (z *WebHandler) Setup(g *gin.Engine) {
	g.GET(webPathLogin+"/:user_hash", z.Login)
	g.GET(webPathForbidden, z.Forbidden)
	g.GET(webPathHome, z.Home)
	g.GET("/", z.Instruction)
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

func (z *WebHandler) Forbidden(g *gin.Context) {
	g.String(http.StatusForbidden, z.Kitchen.UI().Text("web.error.forbidden"))
}

func (z *WebHandler) Instruction(g *gin.Context) {
	g.Redirect(http.StatusTemporaryRedirect, "https://github.com/watermint/toolbox")
}

func (z *WebHandler) Home(g *gin.Context) {
	hash, err := g.Cookie(webUserHashCookie)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
	}
	repo, err := app_user.SingleUserRepository(z.Kitchen.Control().Workspace())
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathServerError)
	}
	user, err := repo.Resolve(hash)
	if err != nil {
		g.Redirect(http.StatusTemporaryRedirect, webPathForbidden)
	}

	g.String(http.StatusOK,
		z.Kitchen.UI().Text("web.home.message", app_msg.P("User", user.UserHash())),
	)
}
