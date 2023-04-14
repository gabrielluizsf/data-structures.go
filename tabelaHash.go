package datastructuresgo

import (
    "fmt"
    "hash/fnv"
)

type MyTable struct {
    table map[int][]*KeyValue
}

type KeyValue struct {
    key   string
    value interface{}
}

func NewTable() *MyTable {
    return &MyTable{table: make(map[int][]*KeyValue)}
}

func (ht *MyTable) Put(key string, value interface{}) {
    hash := hash(key)
    if ht.table[hash] == nil {
        ht.table[hash] = make([]*KeyValue, 0)
    }
    ht.table[hash] = append(ht.table[hash], &KeyValue{key: key, value: value})
}

func (ht *MyTable) Get(key string) interface{} {
    hash := hash(key)
    if ht.table[hash] != nil {
        for _, kv := range ht.table[hash] {
            if kv.key == key {
                return kv.value
            }
        }
    }
    return nil
}

func hash(key string) int {
    h := fnv.New32a()
    h.Write([]byte(key))
    return int(h.Sum32() % 100000)
}

func HashTable() {
	PrintMessage("\nTabela Hash:")
    ht := NewTable()

    // Adicionando algumas entradas à tabela hash
    ht.Put("key1", "comer")
    ht.Put("key2", "dormir")
    ht.Put("key3", "banhar")
	ht.Put("key4", "trabalhar")
	ht.Put("key5", "procrastinar")
	ht.Put("key6", "fazer exercício")
	ht.Put("key7", "ir para a igreja")

    // Obtendo valores a partir de chaves
    fmt.Println("Valor da Key 1: ",ht.Get("key1"))
    fmt.Println("Valor da Key 2: ",ht.Get("key2"))
    fmt.Println("Valor da Key 3: ",ht.Get("key3"))
	fmt.Println("Valor da Key 4: ",ht.Get("key4"))
	fmt.Println("Valor da Key 5: ",ht.Get("key5"))
}
