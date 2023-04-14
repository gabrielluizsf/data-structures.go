package datastructuresgo

import (
    "fmt"
)

type Heap []int

func NewHeap() *Heap {
    heap := Heap{}
    return &heap
}

func (h *Heap) Len() int {
    return len(*h)
}

func (h *Heap) Insert(value int) {
    *h = append(*h, value)
    h.siftUp(len(*h) - 1)
}

func (h *Heap) Extract() (int, bool) {
    if len(*h) == 0 {
        return 0, false
    }
    value := (*h)[0]
    *h = (*h)[1:]
    if len(*h) > 0 {
        h.siftDown(0)
    }
    return value, true
}

func (h *Heap) siftUp(i int) {
    for i > 0 {
        parent := (i - 1) / 2
        if (*h)[parent] >= (*h)[i] {
            break
        }
        (*h)[parent], (*h)[i] = (*h)[i], (*h)[parent]
        i = parent
    }
}

func (h *Heap) siftDown(i int) {
    for {
        left := 2*i + 1
        right := 2*i + 2
        largest := i
        if left < len(*h) && (*h)[left] > (*h)[largest] {
            largest = left
        }
        if right < len(*h) && (*h)[right] > (*h)[largest] {
            largest = right
        }
        if largest == i {
            break
        }
        (*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
        i = largest
    }
}

func Heaps() {

    PrintMessage("\nHeap:")

    heap := NewHeap()
    heap.Insert(2)
    heap.Insert(5)
    heap.Insert(3)
    heap.Insert(1)
    heap.Insert(6)
    heap.Insert(8)
    heap.Insert(4)
    heap.Insert(7)

    for len(*heap) > 0 {
        value, _ := heap.Extract()
        fmt.Println(value)
    }
}
