package datastructuresgo

import (
	"container/list"
	"fmt"
)

func Queue() {
	/*
		cria uma fila utilizando a função New do pacote list
		func New() *List { return new(List).Init() }
		type List struct {
			root Element // sentinel list element, only &root, root.prev, and root.next are used
			len  int     // current list length excluding (this) sentinel element
		}
	*/
	queue := list.New()

	// Adiciona elementos à fila
	queue.PushBack("Alice")
	queue.PushBack("Bob")
	queue.PushBack("Charlie")

	PrintMessage("\nFila:")

	// Percorre a fila e imprimir seus elementos
	for queue.Len() > 0 {
		// Obtem o próximo elemento da fila
		elem := queue.Front()

		// Imprime o elemento
		fmt.Println(elem.Value)

		// Remove o elemento da fila
		queue.Remove(elem)
	}
}
