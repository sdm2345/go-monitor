package monitor

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"log"
	"time"
)

type Conf struct {
	Path                  string
	EnableHttpRequestTime bool
	EnableUptimeSecond    bool
}

func NewDefaultConf() *Conf {
	return &Conf{
		Path:                  "/metrics",
		EnableHttpRequestTime: true,
		EnableUptimeSecond:    true,
	}
}

var startTime = time.Now()
var HttpRequestTime = prometheus.NewHistogram(prometheus.HistogramOpts{
	Name:    "http_request_time",
	Help:    "http_request_time",                  // Sorry, we can't measure how badly it smells.
	Buckets: prometheus.LinearBuckets(0, 0.5, 10), // 5 buckets, each 5 centigrade wide.
})

func InitPrometheus(conf *Conf) {

	if conf.EnableHttpRequestTime {
		err := prometheus.Register(HttpRequestTime)
		if err != nil {
			log.Println("error", err)
		}
	}

	if conf.EnableUptimeSecond {
		promauto.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "process_uptime_seconds",
			Help: "process_uptime_seconds",
		}, func() float64 {
			return time.Since(startTime).Seconds()
		})
	}
}
