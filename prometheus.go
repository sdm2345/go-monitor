package go_monitor

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"log"
	"time"
)

type MonitorConf struct {
	Path string
}

func RegisterGin(r *gin.Engine, conf MonitorConf) {
	p := ginprometheus.NewPrometheus("gin")
	p.MetricsPath = conf.Path
	r.Use(ginMiddleware)
	p.Use(r)

}

var startTime = time.Now()
var httpRequestTime = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "http_request_time",
	Help:    "http_request_time",                  // Sorry, we can't measure how badly it smells.
	Buckets: prometheus.LinearBuckets(0, 0.5, 10), // 5 buckets, each 5 centigrade wide.
})

func ginMiddleware(c *gin.Context) {
	t := time.Now()
	c.Next()
	t2 := time.Now()
	diff := t2.Sub(t).Seconds()
	log.Println("load request", diff)
	httpRequestTime.Observe(diff)
}

func init() {

	promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "process_uptime_seconds",
		Help: "process_uptime_seconds",
	}, func() float64 {
		return time.Since(startTime).Seconds()
	})

	prometheus.Register(httpRequestTime)

}
