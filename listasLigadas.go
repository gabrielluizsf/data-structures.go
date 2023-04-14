package datastructuresgo

import "fmt"

type NodeList struct {
    value int
    next  *NodeList
}

type LinkedList struct {
    head *NodeList
}

func (list *LinkedList) Add(value int) {
    newNodeL := &NodeList{value: value, next: nil}

    if list.head == nil {
        list.head = newNodeL
    } else {
        current := list.head
        for current.next != nil {
            current = current.next
        }
        current.next = newNodeL
    }
}

func (list *LinkedList) Remove(value int) {
    if list.head == nil {
        return
    }

    if list.head.value == value {
        list.head = list.head.next
        return
    }

    current := list.head
    for current.next != nil && current.next.value != value {
        current = current.next
    }
    if current.next != nil {
        current.next = current.next.next
    }
}

func (list *LinkedList) Print() {
    current := list.head
    for current != nil {
        fmt.Printf("%d -> ", current.value)
        current = current.next
    }
    fmt.Println("nil")
}

func LinkedLists() {

	PrintMessage("\nListas Ligadas:")

    list := LinkedList{}

    list.Add(1)
    list.Add(2)
    list.Add(3)

    list.Print() // Output: 1 -> 2 -> 3 -> nil

    list.Remove(3)
	list.Remove(1)

    list.Print() // Output: 2 -> nil
}
