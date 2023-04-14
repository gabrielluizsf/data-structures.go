package datastructuresgo

import "fmt"

type StackOfItems struct {
    items []int
}

func (s *StackOfItems) Push(item int) {
    s.items = append(s.items, item)
}

func (s *StackOfItems) Pop() int {
    if len(s.items) == 0 {
        panic("Pilha vazia!")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item
}

func (s *StackOfItems)printStack() {
	PrintMessage("\nPilha:")
	fmt.Println(s);
}

func Stack() {

    // Cria uma nova pilha
    s := StackOfItems{}

    // Adiciona elementos à pilha
    s.Push(0)
    s.Push(1)
    s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push(6)
	s.Push(7)

	// Imprime os elementos da pilha
	s.printStack()

    // Remove elementos da pilha
	PrintMessage("\nElementos removidos da pilha")
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())
    fmt.Println(s.Pop())

	//Imprime os elementos da pilha
	s.printStack()
}
