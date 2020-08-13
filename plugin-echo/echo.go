package plugin_echo

import (
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sdm2345/go-monitor/monitor"
	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
	"github.com/slok/go-http-metrics/middleware"
	echomiddleware "github.com/slok/go-http-metrics/middleware/echo"
)

func Init(r *echo.Echo, options ...monitor.Option) {

	conf := monitor.NewDefaultConf(options)

	m := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})

	r.Use(echomiddleware.Handler("", m))
	r.GET(conf.Path, echo.WrapHandler(promhttp.Handler()))

}
