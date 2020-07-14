package plugin_gin

import (
	"github.com/gin-gonic/gin"
	"github.com/sdm2345/go-monitor/monitor"
	"github.com/zsais/go-gin-prometheus"
	"time"
)

type Option func(conf *monitor.Conf)

func WithPath(path string) Option {
	return func(conf *monitor.Conf) {
		conf.Path = path
	}
}

func RegisterGin(r *gin.Engine, options ...Option) {
	p := ginprometheus.NewPrometheus("gin")
	conf := monitor.NewDefaultConf()
	for _, op := range options {
		op(conf)
	}
	p.MetricsPath = conf.Path

	monitor.InitPrometheus(conf)
	if conf.EnableHttpRequestTime {
		r.Use(ginMiddleware)
	}
	p.Use(r)

}

func ginMiddleware(c *gin.Context) {
	t := time.Now()
	c.Next()
	diff := time.Now().Sub(t).Seconds()
	monitor.HttpRequestTime.Observe(diff)
}
