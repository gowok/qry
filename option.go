package qry

type option struct {
	Driver string
}

type OptionFunc func(*option)

func WithDriver(driver string) OptionFunc {
	return func(o *option) {
		o.Driver = driver
	}
}

func NewOption() *option {
	return &option{}
}
