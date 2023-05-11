package qry

type option struct {
	driver string
}

type OptionFunc func(*option)

func WithDriver(driver string) OptionFunc {
	return func(o *option) {
		o.driver = driver
	}
}
