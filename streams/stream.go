package stream

type Stream[T any] struct {
	items      []T
	operations []Operation[T]
	parallel   bool
	ordered    bool
}

func Of[T any](items []T) *Stream[T] {
	return &Stream[T]{items: items, ordered: true}
}

func (s *Stream[T]) Parallel() *Stream[T] {
	s.parallel = true
	return s
}

func (s *Stream[T]) Unordered() *Stream[T] {
	s.ordered = false
	return s
}
