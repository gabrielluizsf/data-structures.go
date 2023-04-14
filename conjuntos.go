package datastructuresgo

import "fmt"

type Set map[string]bool

func (s Set) Add(element string) {
    s[element] = true
}

func (s Set) Remove(element string) {
    delete(s, element)
}

func (s Set) Contains(element string) bool {
    return s[element]
}

func (s Set) Size() int {
    return len(s)
}

func Sets() {
    PrintMessage("\nConjuntos:")
    set := make(Set)
    set.Add("element1")
    set.Add("element2")
    set.Add("element3")
    fmt.Println("Existe o Elemento 2?",set.Contains("element2")) // true
    fmt.Println("Existe o Elemento 4?",set.Contains("element4")) // false
    set.Add("element4")
    fmt.Println("Existe o Elemento 4?",set.Contains("element4")) // false
    set.Remove("element1")
    set.Remove("element3")
    fmt.Println("Quantidade de Elementos da lista:",set.Size()) // 2
}
