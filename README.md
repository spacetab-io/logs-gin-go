logs-go-gin
-----------

Logger fo gin gonic

## Usage

Initiate new [logger](https://github.com/microparts/logs-go) with prefilled `logs.Config` and use it to initiate new gin logger

```go
package main

import (
	"time"
	
  "github.com/gin-gonic/gin"
	"github.com/microparts/logs-go"
	"github.com/microparts/logs-go-gin"
)

func main() {
	// initiate logs-go
	conf := &logs.Config{
		Level:"warn",
		Format: "text",
		Sentry: &logs.SentryConfig{
			Enable: true,
			Stage:"test",
			DSN: "http://dsn.sentry.com",
			ResponseTimeout: 0,
			StackTrace: logs.StackTraceConfig{
				Enable: true,
			},
		},
	}
	
	l, err := logs.NewLogger(conf)
	if err != nil {
		panic(err)
	}
	
	r := gin.New()
  r.Use(ginLogger.NewLogger(l), gin.Recovery())
  
  // pingpong
  r.GET("/ping", func(c *gin.Context) {
  	c.Data(200, "text/plain", []byte("pong"))
  })
  
  _ = r.Run("127.0.0.1:8080")
}
```

## Licence

The software is provided under [MIT Licence](LICENCE).