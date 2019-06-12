package kite_base

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/experimental/app_root"
	"time"
)

func Start() {
	l := app_root.Log()

	g := gin.New()
	g.Use(ginzap.Ginzap(l, time.RFC3339, true))
	g.Use(ginzap.RecoveryWithZap(l, true))

	_ = g.Run(":7800")
}
