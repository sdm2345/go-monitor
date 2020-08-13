package plugin_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdm2345/go-monitor/monitor"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	ginmiddleware "github.com/slok/go-http-metrics/middleware/gin"
)

func Init(r *gin.Engine, options ...monitor.Option) {

	conf := monitor.NewDefaultConf(options)

	m := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r.Use(ginmiddleware.Handler("", m))
	r.GET(conf.Path, gin.WrapH(promhttp.Handler()))

}
