package datastructuresgo

import "fmt"

func Arrays() {
	PrintMessage("\nArrays:")
    // Cria um array com espaço para 5 elementos
    var a [5]int

    // Atribue valores aos elementos do array
    a[0] = 1
    a[1] = 2
    a[2] = 3
    a[3] = 4
    a[4] = 5

    // Imprime o array
    fmt.Println(a)

    // Cria um array com valores iniciais
    b := [3]int{1, 2, 3}

    // Imprime o array
    fmt.Println(b)

    // Obtem o tamanho do array
    fmt.Println("Tamanho do array 'a':", len(a))
    fmt.Println("Tamanho do array 'b':", len(b))

    // Acessa um elemento do array através do índice
    fmt.Println("Elemento na posição 2 do array 'a':", a[2])
    fmt.Println("Elemento na posição 2 do array 'b':", b[2])

    // Altera um elemento do array através do índice
    a[2] = 10
    fmt.Println("Valor do elemento na posição 2 do array 'a' atualizado com sucesso para o valor:", a[2])
    b[2] = 25
    fmt.Println("Valor do elemento na posição 2 do array 'b' atualizado com sucesso para o valor:", b[2])
}
