package main

import (
	"github.com/gin-gonic/gin"

	"github.com/sdm2345/go-monitor/plugin-gin"

	"math/rand"
	"net/http"
	"time"
)

///Users/i538105/code/sap/Eureka/dev-git-monitor-libs/example/go.mod:8:
func main() {
	r := gin.Default()

	plugin_gin.RegisterGin(r, plugin_gin.WithPath("/metrics"))

	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * time.Duration(rand.Int31n(3)))
		c.String(http.StatusOK, "hello world")
	})

	r.Run()
}
