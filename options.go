package options

type Option[T any] interface {
	apply(*T)
}

type Func[T any] func(*T)

func (f Func[T]) apply(options *T) {
	f(options)
}

func New[T any](opts ...Option[T]) *T {
	options := new(T)

	Apply(options, opts...)

	return options
}

func Apply[T any](options *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt.apply(options)
	}
}
