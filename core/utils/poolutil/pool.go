package poolutil

type Pool[T any] interface {
	Get() T
	Put(T)
}
