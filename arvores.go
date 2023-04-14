package datastructuresgo

import "fmt"

type Person struct {
    Name     string
    Age      int
    Children []*Person
}

func (p *Person) AddChild(child *Person) {
    p.Children = append(p.Children, child)
}

func (p *Person) PrintTree() {
    fmt.Println(p.Name)
    for _, child := range p.Children {
        child.PrintTree()
    }
}

func (p *Person) NumberOfChildren() int{
	number := 0;
	for _,  child := range p.Children {
	 child.NumberOfChildren()
	 return	number + 1;
	}
	return number;
}

func(p *Person) PrintFamily(){
	fmt.Printf("\nNome:");
	fmt.Println(p.Name,"\nIdade:",p.Age)
	PrintMessage("\nFilhos:")
	for i := 0; i <=  p.NumberOfChildren(); i++{
		fmt.Println(i + 1,"º Filho:",p.Children[i].Name,"\tIdade:",p.Children[i].Age)
	}
}

func Trees() {

	PrintMessage("\nÁRVORES:")

    // Cria algumas pessoas
    alice := &Person{Name: "Alice", Age: 70}
    bob := &Person{Name: "Bob", Age: 55}
    charlie := &Person{Name: "Charlie", Age: 25}
    dave := &Person{Name: "Dave", Age: 10}
    eve := &Person{Name: "Eve", Age: 10}
    frank := &Person{Name: "Frank", Age: 5}
	angelica := &Person{Name: "Ângelica", Age: 7}

    // Adiciona filhos à árvore genealógica
    alice.AddChild(bob)
    alice.AddChild(charlie)
    bob.AddChild(dave)
    bob.AddChild(eve)
    charlie.AddChild(frank)
	charlie.AddChild(angelica)

    // Imprime a árvore genealógica
    alice.PrintTree()

	alice.PrintFamily()
	bob.PrintFamily()
	charlie.PrintFamily()
	
}
