package ut_log

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

// Modified version of Ginzap func
func GinWrapper(l *zap.Logger) gin.HandlerFunc {
	return func(g *gin.Context) {
		path := g.Request.URL.Path
		query := g.Request.URL.RawQuery

		ll := l.With(zap.String("path", path), zap.String("query", query))

		start := time.Now()
		g.Next()
		end := time.Now()
		latency := end.Sub(start)

		switch {
		case 0 < len(g.Errors):
			for _, err := range g.Errors.Errors() {
				ll.Debug("error", zap.String("error", err))
			}
		default:
			ll.Debug(
				g.Request.Method,
				zap.Int("status", g.Writer.Status()),
				zap.String("ip", g.ClientIP()),
				zap.String("user_agent", g.Request.UserAgent()),
				zap.String("time", end.Format(time.RFC3339)),
				zap.String("latency", latency.String()),
			)
		}
	}
}
