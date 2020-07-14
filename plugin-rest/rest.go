package plugin_rest

import (
	"github.com/emicklei/go-restful"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdm2345/go-monitor/monitor"
	"time"
)

type Option func(conf *monitor.Conf)

func WithPath(path string) Option {
	return func(conf *monitor.Conf) {
		conf.Path = path
	}
}

func RegisterRest(c *restful.Container, options ...Option) {
	conf := monitor.NewDefaultConf()
	for _, op := range options {
		op(conf)
	}
	monitor.InitPrometheus(conf)
	c.Handle(conf.Path, promhttp.Handler())
	if conf.EnableHttpRequestTime {
		c.Filter(func(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {

			t1 := time.Now()
			chain.ProcessFilter(request, response)
			t2 := time.Now()
			monitor.HttpRequestTime.Observe(t2.Sub(t1).Seconds())
		})
	}

}
