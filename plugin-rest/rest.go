package plugin_rest

import (
	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdm2345/go-monitor/monitor"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	gorestfulmiddleware "github.com/slok/go-http-metrics/middleware/gorestful"
)

func Init(c *restful.Container, options ...monitor.Option) {
	conf := monitor.NewDefaultConf(options)

	m := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	c.Filter(gorestfulmiddleware.Handler("", m))
	c.Handle(conf.Path, promhttp.Handler())

}
