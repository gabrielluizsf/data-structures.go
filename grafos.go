package datastructuresgo

import (
	"fmt"
)

type Graph struct {
    Nodes []*Node
}

type Node struct {
    Name     string
    Edges    []*Edge
    Visited  bool
}

type Edge struct {
    FromNode *Node
    ToNode   *Node
    Weight   int
}

func (g *Graph) AddNode(n *Node) {
    g.Nodes = append(g.Nodes, n)
}

func (n *Node) AddEdge(e *Edge) {
    n.Edges = append(n.Edges, e)
}

func (g *Graph) PrintGraph() {
    for _, n := range g.Nodes {
        fmt.Printf("%s -> ", n.Name)
        for _, e := range n.Edges {
            fmt.Printf("%s (%d), ", e.ToNode.Name, e.Weight)
        }
        fmt.Println("")
    }
}


func GraphAlphabet() {
	PrintMessage("\nGrafos:")
    // Cria alguns nós
    a := &Node{Name: "A"}
    b := &Node{Name: "B"}
    c := &Node{Name: "C"}
    d := &Node{Name: "D"}
    e := &Node{Name: "E"}
	f := &Node{Name: "F"}
	g := &Node{Name: "G"}
	h := &Node{Name: "H"}
	i := &Node{Name: "I"}
	j := &Node{Name: "J"}
	k := &Node{Name: "K"}
	l := &Node{Name: "L"}
	m := &Node{Name: "M"}
	n := &Node{Name: "N"} 
	o := &Node{Name: "O"} 
	p := &Node{Name: "P"} 
	q := &Node{Name: "Q"}
	r := &Node{Name: "R"}
	s := &Node{Name: "S"}
	t := &Node{Name: "T"}
	u := &Node{Name: "U"}
	v := &Node{Name: "V"}
	w := &Node{Name: "W"}
	x := &Node{Name: "X"}
	y := &Node{Name: "Y"}
	z := &Node{Name: "z"}
	// Cria algumas arestas


    // Adiciona os nós e arestas ao grafo
    graph := &Graph{}
    graph.AddNode(a)
    graph.AddNode(b)
    graph.AddNode(c)
    graph.AddNode(d)
    graph.AddNode(e)
	graph.AddNode(f)
	graph.AddNode(g)
	graph.AddNode(h)
	graph.AddNode(i)
	graph.AddNode(j)
	graph.AddNode(k)
	graph.AddNode(l)
	graph.AddNode(m)
	graph.AddNode(n)
	graph.AddNode(o)
	graph.AddNode(p)
	graph.AddNode(q)
	graph.AddNode(r)
	graph.AddNode(s)
	graph.AddNode(t)
	graph.AddNode(u)
	graph.AddNode(v)
	graph.AddNode(w)
	graph.AddNode(x)
	graph.AddNode(y)
	graph.AddNode(z)
	
	ab := &Edge{FromNode: a, ToNode: b, Weight: 1}
	ay := &Edge{FromNode: a, ToNode: y, Weight: 8} 
    ac := &Edge{FromNode: a, ToNode: c, Weight: 3}
	az := &Edge{FromNode: a, ToNode: z, Weight: 12}
    bc := &Edge{FromNode: b, ToNode: c, Weight: 2}
    bd := &Edge{FromNode: b, ToNode: d, Weight: 5}
	by := &Edge{FromNode: b, ToNode: y,Weight: 9}
    cd := &Edge{FromNode: c, ToNode: d, Weight: 4}
    ce := &Edge{FromNode: c, ToNode: e, Weight: 6}
    de := &Edge{FromNode: d, ToNode: e, Weight: 7}
	ya := &Edge{FromNode: y, ToNode: a, Weight: 13}
	yz := &Edge{FromNode: y, ToNode: z, Weight: 15}
	yb := &Edge{FromNode: y, ToNode: b, Weight: 14}
	za := &Edge{FromNode: z, ToNode: a, Weight: 11}
	zy := &Edge{FromNode: z, ToNode: y, Weight: 10}
    a.AddEdge(ab)
    a.AddEdge(ac)
	a.AddEdge(ay)
	a.AddEdge(az)
    b.AddEdge(bc)
    b.AddEdge(bd)
	b.AddEdge(by)
    c.AddEdge(cd)
    c.AddEdge(ce)
    d.AddEdge(de)
	y.AddEdge(yz)
	y.AddEdge(ya)
	y.AddEdge(yb)
	z.AddEdge(za)
	z.AddEdge(zy)


    // Imprime o grafo
    graph.PrintGraph()
}
