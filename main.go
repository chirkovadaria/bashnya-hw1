package main

import "fmt"

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{
		items: make([]int, 0),
	}
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	lastItem := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return lastItem
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	size := len(s.items)
	return size
}

func (s *Stack) Clear() []int {
	s.items = s.items[:0]
	return s.items
}

func main() {
	stack := NewStack()

	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Удаление и возврат последнего элемента из стека:")
	fmt.Println(stack.Pop())

	fmt.Println("Проверка пустой стек или нет:")
	fmt.Println(stack.IsEmpty())

	fmt.Println("Размер стека:")
	fmt.Println(stack.Size())

	fmt.Println("Очистка стека:")
	fmt.Println(stack.Clear())

}
