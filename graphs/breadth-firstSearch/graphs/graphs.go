package graphs

import (
	"algorithms/graphs/breadth-firstSearch/singlyLinkedList"
	"fmt"
)

const (
	Gray="gray"
	White="white"
	Black="black"
)

type Graph struct{
	Nodes []Node
}

type Node struct{
	Nvertex int //номер вершины,
	Color string
	List singlyLinkedList.List
	D int//расстояние до исходной вершины
	Parent int
}

func (G *Graph) InsertNode(Nvertex int,List singlyLinkedList.List){
	G.Nodes=append(G.Nodes,Node{
		Nvertex:Nvertex,
		List:List,
	})
}
func (G *Graph) Display(){
	for _,node := range G.Nodes{
		fmt.Println("Vertex=",node.Nvertex)
		node.List.Display()
	}
}


