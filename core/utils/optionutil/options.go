package optionutil

type I[T any] interface {
	apply(*T)
}

var _ I[any] = (*F[any])(nil)

type F[T any] struct {
	f func(*T)
}

//nolint:unused
func (f *F[T]) apply(opts *T) {
	f.f(opts)
}

func New[T any](f func(*T)) *F[T] {
	return &F[T]{f}
}

func Build[T any](def *T, opts ...I[T]) *T {
	for i := 0; i < len(opts); i++ {
		opts[i].apply(def)
	}
	return def
}
