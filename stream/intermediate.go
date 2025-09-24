package stream

func (s *Stream[T]) Filter(f func(T) bool) *Stream[T] {
	s.operations = append(s.operations, FilterOperation[T]{f: f})
	return s
}

func (s *Stream[T]) Map(f func(T) T) *Stream[T] {
	s.operations = append(s.operations, MapOperation[T]{f: f})
	return s
}
