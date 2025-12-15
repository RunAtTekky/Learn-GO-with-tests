package generics

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Push(element T) {
	s.values = append(s.values, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	last_idx := len(s.values) - 1
	element := s.values[last_idx]

	s.values = s.values[:last_idx]

	return element, true
}
