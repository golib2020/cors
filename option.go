package cors

type options struct {
	origin  string
	methods []string
	headers []string
}

type Option func(*options)

func WithOrigin(origin string) Option {
	return func(opts *options) {
		opts.origin = origin
	}
}

func WithMethods(methods ...string) Option {
	return func(opts *options) {
		opts.methods = methods
	}
}

func WithHeaders(headers ...string) Option {
	return func(opts *options) {
		opts.headers = headers
	}
}
