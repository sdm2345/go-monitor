package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sdm2345/go-monitor/monitor"
	"github.com/sdm2345/go-monitor/plugin-gin"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()

	plugin_gin.Init(r, monitor.WithPath("/metrics"))

	r.GET("/", func(c *gin.Context) {
		time.Sleep(time.Second * time.Duration(rand.Int31n(3)))
		c.String(http.StatusOK, "hello world")
	})
	r.GET("/500", func(c *gin.Context) {

		c.String(http.StatusOK,
			fmt.Sprint("hello world:%d", 2/rand.Intn(2)))
	})

	r.Run(":8094")
}
