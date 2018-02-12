package graphs

import "algorithms/graphs/breadth-firstSearch/singlyLinkedList"

type Graph struct{
	lists []singlyLinkedList.List
}

func (G *Graph) InsertList(L singlyLinkedList.List){
	G.lists=append(G.lists,L)
}
func (G *Graph) Display(){
	for _,l := range G.lists{
		l.Display()
	}
}

