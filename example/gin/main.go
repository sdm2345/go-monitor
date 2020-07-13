package main

import (
	"github.com/gin-gonic/gin"
	monitor "github.wdf.sap.corp/Eureka/dev-git-monitor-libs"
	"math/rand"
	"net/http"
	"time"
)

///Users/i538105/code/sap/Eureka/dev-git-monitor-libs/example/go.mod:8:
func main() {
	r := gin.Default()

	monitor.RegisterGin(r, monitor.MonitorConf{
		Path: "/metrics",
	})

	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second*time.Duration(rand.Int31n(3)))
		c.String(http.StatusOK, "hello world")
	})

	r.Run()
}
