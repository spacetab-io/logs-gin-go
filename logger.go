package ginLogger

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// NewLogger logs gin gonic actions
func NewLogger(l *logrus.Logger) gin.HandlerFunc {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	var skip map[string]struct{}
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		if _, ok := skip[path]; !ok {
			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()

			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			if raw != "" {
				path = path + "?" + raw
			}

			// Stop timer
			end := time.Now()
			latency := end.Sub(start)
			dataLength := c.Writer.Size()
			if dataLength < 0 {
				dataLength = 0
			}

			log := l.WithFields(logrus.Fields{
				"hostname":   hostname,
				"latency":    latency,
				"clientIP":   clientIP,
				"status":     statusCode,
				"proto":      c.Request.Proto,
				"method":     method,
				"path":       path,
				"referer":    c.Request.Referer(),
				"query":      raw,
				"comment":    comment,
				"dataLength": dataLength,
			})

			if len(c.Errors) > 0 {
				log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				if statusCode > 499 {
					log.Error("http request error")
				} else if statusCode > 399 {
					log.Warn("http request warning")
				} else {
					log.Info("http request info")
				}
			}
		}
	}
}
