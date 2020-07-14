package monitor

type Conf struct {
	Path string
}

func NewDefaultConf(options []Option) *Conf {
	conf := &Conf{
		Path: "/metrics",
	}
	for _, op := range options {
		op(conf)
	}
	return conf
}

type Option func(conf *Conf)

func WithPath(path string) Option {
	return func(conf *Conf) {
		conf.Path = path
	}
}

func init() {

	InitUptime()
}
