package stream

type Operation[T any] interface {
	Apply(v T) (T, bool)
	IsOrdered() bool
}

// FilterOperation filters the stream items.
type FilterOperation[T any] struct {
	f func(T) bool
}

// Apply implements Operation.
func (op FilterOperation[T]) Apply(v T) (T, bool) {
	return v, op.f(v)
}

func (op FilterOperation[T]) IsOrdered() bool {
	return true
}
