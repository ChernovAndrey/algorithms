package breadth_firstSearch

import (
	"algorithms/graphs/breadth-firstSearch/graphs"
	"algorithms/graphs/breadth-firstSearch/queue"
	"fmt"
)

func Search(G *graphs.Graph, s int) { //s- номер исходной вершины
	curNode := &graphs.Node{}
	for i:=0;i<len(G.Nodes);i++ {
		G.Nodes[i].Color = graphs.White
		G.Nodes[i].D = -1 // если -1 то недостежима
		G.Nodes[i].Parent = -1 // нет родителя
		if G.Nodes[i].Nvertex == s {
			curNode = &G.Nodes[i]
		}
	}

	curNode.Color = graphs.Gray
	curNode.D = 0
	curNode.Parent = -1 //нет родителя

	Q := queue.Queue{}
	Q.Init()
	Q.EnQueue(curNode)
	fmt.Println(Q.IsEmpty())
	for ; !Q.IsEmpty(); {
		curNode = Q.DeQueue().(*graphs.Node)
		NodeList := curNode.List.Head

		for ; NodeList != nil; { // в G.Nodes на нулевом месте эл-т с nVertex=1, на первом nVertex=2 и т д
			nCurVertex := NodeList.Key
			nodeGraph := &G.Nodes[nCurVertex-1]

			if nodeGraph.Color == graphs.White {
				nodeGraph.Color = graphs.Gray
				nodeGraph.D = curNode.D + 1
				nodeGraph.Parent = curNode.Nvertex
				Q.EnQueue(nodeGraph)
			}

			NodeList = NodeList.Next
		}

		curNode.Color=graphs.Black
	}
}

func DistanceFromGivenVertex(G *graphs.Graph,s int){
	Search(G,s)
	for _,node := range G.Nodes{
		fmt.Print("nVertex=",node.Nvertex,"	")
		fmt.Println("Distance=",node.D)
		node.List.Display()
	}
}

// выводит кратчайший путь от s к v
func PrintPath(G *graphs.Graph,s int,v int) { //предполагается что поиск в ширину был уже выполнен
	if s ==v{
		fmt.Print(s,"	")
		return
	}
	parentV:=G.Nodes[v-1].Parent
	if parentV==-1{
		fmt.Println("пути нет")
	}else {
		PrintPath(G,s,parentV)
	}
	fmt.Print(v,"	")
}