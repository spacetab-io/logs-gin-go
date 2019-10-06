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
		c.Next()
		if _, ok := skip[path]; !ok {
			statusCode := c.Writer.Status()
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()
			raw := c.Request.URL.RawQuery
			if raw != "" {
				path = path + "?" + raw
			}

			dataLength := c.Writer.Size()
			if dataLength < 0 {
				dataLength = 0
			}

			log := l.WithFields(logrus.Fields{
				"hostname":    hostname,
				"clientIP":    c.ClientIP(),
				"proto":       c.Request.Proto,
				"method":      c.Request.Method,
				"contentType": c.ContentType(),
				"dataLength":  dataLength,
				"status":      statusCode,
				"path":        path,
				"referer":     c.Request.Referer(),
				"query":       raw,
			})

			if comment != "" {
				log = log.WithField("comment", comment)
			}

			log = log.WithField("latency", time.Now().Sub(start))

			if len(c.Errors) > 0 {
				log.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				if statusCode > 499 {
					log.Error("http request error")
				} else if statusCode > 399 {
					log.Warn("http request warning")
				} else {
					log.Debug("http request info")
				}
			}
		}
	}
}
