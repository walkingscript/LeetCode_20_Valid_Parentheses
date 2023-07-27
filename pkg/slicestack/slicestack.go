package slicestack

type IStack[T any] interface {
	Push(value T)
	Pop() (value T)
	PopLeft() (val T)
	IsEmpty() bool
	Count() int
}

type Stack[T any] struct {
	values []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		values: make([]T, 0, 10),
	}
}

func (s *Stack[T]) Push(val T) {
	s.values = append(s.values, val)
}

func (s *Stack[T]) Pop() (val T) {
	val = s.values[s.Count()-1]
	s.values = s.values[:s.Count()-1]
	return
}

func (s *Stack[T]) PopLeft() (val T) {
	val = s.values[0]
	s.values = s.values[1:]
	return
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Count() == 0
}

func (s *Stack[T]) Count() int {
	return len(s.values)
}
