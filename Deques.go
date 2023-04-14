package datastructuresgo

import "fmt"

type node struct {
    value      interface{}
    prev, next *node
}

type Deque struct {
    head, tail *node
    size       int
}

func NewDeque() *Deque {
    head := &node{}
    tail := &node{}
    head.next = tail
    tail.prev = head
    return &Deque{head, tail, 0}
}

func (d *Deque) Len() int {
    return d.size
}

func (d *Deque) PushFront(v interface{}) {
    n := &node{value: v, prev: d.head, next: d.head.next}
    d.head.next.prev = n
    d.head.next = n
    d.size++
}

func (d *Deque) PushBack(v interface{}) {
    n := &node{value: v, prev: d.tail.prev, next: d.tail}
    d.tail.prev.next = n
    d.tail.prev = n
    d.size++
}

func (d *Deque) PopFront() interface{} {
    if d.size == 0 {
        return nil
    }
    n := d.head.next
    d.head.next = n.next
    n.next.prev = d.head
    d.size--
    return n.value
}

func (d *Deque) PopBack() interface{} {
    if d.size == 0 {
        return nil
    }
    n := d.tail.prev
    d.tail.prev = n.prev
    n.prev.next = d.tail
    d.size--
    return n.value
}

func (d *Deque) Front() interface{} {
    if d.size == 0 {
        return nil
    }
    return d.head.next.value
}

func (d *Deque) Back() interface{} {
    if d.size == 0 {
        return nil
    }
    return d.tail.prev.value
}

func Deques() {

	PrintMessage("\n\nDouble Ended Queue:")

    deque := NewDeque()
    deque.PushBack("apple")
    deque.PushFront("banana")
    deque.PushBack("pear")
    deque.PushFront("orange")

    fmt.Println("Deque size:", deque.Len())
    fmt.Println("Deque front:", deque.Front())
    fmt.Println("Deque back:", deque.Back())

    fmt.Println("Deque contents (front to back):")
    for i := 0; i < deque.Len(); i++ {
        fmt.Println(deque.PopFront())
    }
}
