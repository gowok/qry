package qry

// option holds the settings that refine how a query is generated.
type option struct {
	Driver string
}

// OptionFunc mutates an option set.
type OptionFunc func(*option)

// WithDriver sets the SQL driver a query is generated for.
func WithDriver(driver string) OptionFunc {
	return func(o *option) {
		o.Driver = driver
	}
}

// NewOption returns an empty option set.
func NewOption() *option {
	return &option{}
}
