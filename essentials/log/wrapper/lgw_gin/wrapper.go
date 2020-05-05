package lgw_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// Here is the modified version of gin-contrib/zap
//
// https://github.com/gin-contrib/zap
// License: https://github.com/gin-contrib/zap/blob/master/LICENSE

// Modified version of Ginzap func
// https://github.com/gin-contrib/zap/blob/v0.0.1/zap.go#L65
func GinWrapper(l es_log.Logger) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := g.Request.URL.Path
		query := g.Request.URL.RawQuery

		ll := l.With(es_log.String("path", path), es_log.String("query", query))

		start := time.Now()
		g.Next()
		end := time.Now()
		latency := end.Sub(start)

		switch {
		case 0 < len(g.Errors):
			for _, err := range g.Errors.Errors() {
				ll.Debug("error", es_log.String("error", err))
			}
		default:
			ll.Debug(
				g.Request.Method,
				es_log.Int("status", g.Writer.Status()),
				es_log.String("ip", g.ClientIP()),
				es_log.String("user_agent", g.Request.UserAgent()),
				es_log.String("time", end.Format(time.RFC3339)),
				es_log.String("latency", latency.String()),
			)
		}
	}
}

// Modified version of RecoveryWithZap
// https://github.com/gin-contrib/zap/blob/v0.0.1/zap.go#L65
func ginRecoveryHandler(l es_log.Logger, c *gin.Context, err interface{}) {
	// Check for a broken connection, as it is not really a
	// condition that warrants a panic stack trace.
	var brokenPipe bool
	if ne, ok := err.(*net.OpError); ok {
		if se, ok := ne.Err.(*os.SyscallError); ok {
			if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
				brokenPipe = true
			}
		}
	}

	httpRequest, _ := httputil.DumpRequest(c.Request, false)
	if brokenPipe {
		l.Warn(c.Request.URL.Path,
			es_log.Any("error", err),
			es_log.String("request", string(httpRequest)),
		)
		_ = c.Error(err.(error))
		c.Abort()
		return
	}

	l.Warn("[Recovery from panic]",
		es_log.Time("time", time.Now()),
		es_log.Any("error", err),
		es_log.String("request", string(httpRequest)),
		es_log.String("stack", string(debug.Stack())),
	)

	c.AbortWithStatus(http.StatusInternalServerError)
}

func GinRecovery(l es_log.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ginRecoveryHandler(l, context, err)
			}
			context.Next()
		}()
	}
}
