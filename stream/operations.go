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

// MapOperation transforms the stream items.
type MapOperation[T any] struct {
	f func(T) T
}

// Apply implements Operation.
func (op MapOperation[T]) Apply(v T) (T, bool) {
	return op.f(v), true
}

func (op MapOperation[T]) IsOrdered() bool {
	return true
}

// FlatMapOperation transforms the stream items to a new stream.
type FlatMapOperation[T any, R any] struct {
	f func(T) []R
}

func (op FlatMapOperation[T, R]) Apply(v T) ([]R, bool) {
	return op.f(v), true
}

func (op FlatMapOperation[T, R]) IsOrdered() bool {
	return true
}
