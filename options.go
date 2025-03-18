package options

type Option[T any] interface {
	apply(*T)
}

type optionFunc[T any] func(*T)

func (f optionFunc[T]) apply(options *T) {
	f(options)
}

func Func[T any](f func(*T)) Option[T] {
	return optionFunc[T](f)
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
