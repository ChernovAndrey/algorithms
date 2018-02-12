package singlyLinkedList

import "fmt"

type Node struct {
	next *Node
	key int
}

type List struct{
	head *Node
}

func (L *List) Insert(key int){
	node:= &Node{
		next:L.head,
		key:key,
	}
	L.head=node
}

func (L *List) InsertArray(keys []int){
	for _,k := range keys{
		L.Insert(k)
	}
}

func (L *List) Search(key int) *Node{
	x:=L.head
	for ;(x!=nil)&&(x.key!=key);{
		x=x.next
	}
	return x
}

func (L* List) Delete(key int){
	if L.head.key==key {
		L.head=L.head.next
	}
	x:=L.head
	for ;x.next!=nil;{
		if x.next.key==key{
			x.next=x.next.next
			return
		}
		x=x.next
	}
}

func (L* List) Display(){
	x:=L.head
	for ;x!=nil;{
		fmt.Print(x.key)
		fmt.Print("->")
		x=x.next
	}
	fmt.Println()
}
