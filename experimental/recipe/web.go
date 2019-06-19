package recipe

import (
	"fmt"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_vo"
	"time"
)

const (
	webPort = 7800
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

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))

	_ = g.Run(fmt.Sprintf(":%d", wvo.Port))
	panic("implement me")
}

type WebHandler struct {
	Kitchen app_recipe.Kitchen
}

func (z *WebHandler) Setup(g *gin.Engine) {
	g.GET("/login/:user_hash", z.Login)
}

func (z *WebHandler) Login(g *gin.Context) {

}
