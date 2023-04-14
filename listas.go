package datastructuresgo

import (
    "container/list"
    "fmt"
)

func Lists() {
	PrintMessage("\nLista:")

    // Cria uma nova lista
    l := list.New()

    // Adiciona elementos à lista
    l.PushBack(0)
    l.PushBack(42)
    l.PushBack(11)
    l.PushBack(722)
    l.PushBack(53)
    l.PushBack(100)

    // Percorre a lista e imprimir seus elementos
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }

    // Remove o primeiro elemento da lista
    l.Remove(l.Front())

    // Imprime a lista após a remoção
    fmt.Println("\nLista sem o primeiro Elemento:")
    for e := l.Front(); e != nil; e = e.Next() {
        fmt.Println(e.Value)
    }
}
