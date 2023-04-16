package spec

type (
	Variable[T any] interface {
		Value() T
		Clear()
		ClearCache(func() T)
	}

	variable[T any] struct {
		value  T
		cache  func() T
		cached bool
	}
)

func (it *variable[T]) init(cache func() T) *variable[T] {
	it.cache = cache
	it.cached = false

	return it
}

func (it *variable[T]) Value() T {
	if !it.cached {
		it.value = it.cache()
	}

	return it.value
}

func (it *variable[T]) Clear() {
	it.cached = false
}

func (it *variable[T]) ClearCache(cache func() T) {
	defer it.Clear()

	it.cache = cache
}
