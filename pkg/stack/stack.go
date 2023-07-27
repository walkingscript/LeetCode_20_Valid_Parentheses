package stack

type IStack[T any] interface {
	Push(value T)
	Pop() (value T)
	IsEmpty() bool
}

type Stack[T any] struct {
	node *Item[T]
}

type Item[T any] struct {
	value T
	prev  *Item[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(val T) {
	if s.node == nil {
		s.node = &Item[T]{value: val}
		return
	}
	buf := s.node
	s.node = &Item[T]{value: val}
	s.node.prev = buf
}

func (s *Stack[T]) Pop() (val T) {
	val = s.node.value
	s.node = s.node.prev
	return
}

func (s *Stack[T]) IsEmpty() bool {
	return s.node == nil
}
