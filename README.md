logs-go-gin
-----------

Logger fo gin gonic

## Usage

Pass logrus instance to constructor.

```go
package main

import (
    "time"

    ginLogger "github.com/microparts/logs-go-gin"

    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func main() {
    r := gin.New()
    r.Use(ginLogger.NewLogger(logrus.New()), gin.Recovery())
  
    // pingpong
    r.GET("/ping", func(c *gin.Context) {
        c.Data(200, "text/plain", []byte("pong"))
    })
    _ = r.Run("127.0.0.1:8080")
}
```

## Licence

The software is provided under [MIT Licence](LICENCE).