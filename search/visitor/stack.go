package visitor

type Stack[T any] struct {
	queue []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil}
}

func (stack *Stack[T]) Push(item T) {
	stack.queue = append([]T{item}, stack.queue...)
}

func (stack *Stack[T]) Top() (T, bool) {
	var x T
	if len(stack.queue) > 0 {
		x = stack.queue[len(stack.queue)-1]
		return x, true
	}
	return x, false
}

func (stack *Stack[T]) Pop() T {
	if len(stack.queue) == 0 {
		panic("stack is empty")
	}

	item := stack.queue[len(stack.queue)-1]
	stack.queue = stack.queue[:len(stack.queue)-1]
	return item
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.queue) == 0
}

func (stack *Stack[T]) Len() int {
	return len(stack.queue)
}
