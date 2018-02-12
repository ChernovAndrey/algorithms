package singlyLinkedList

import "fmt"

type node struct {
	Next *node
	Key int
}

type List struct{
	Head *node
}

func (L *List) Insert(key int){
	node:= &node{
		Next:L.Head,
		Key:key,
	}
	L.Head=node
}

func (L *List) InsertArray(Keys []int){
	for _,k := range Keys{
		L.Insert(k)
	}
}

func (L *List) Search(key int) *node{
	x:=L.Head
	for ;(x!=nil)&&(x.Key!=key);{
		x=x.Next
	}
	return x
}

func (L* List) Delete(key int){
	if L.Head.Key==key {
		L.Head=L.Head.Next
	}
	x:=L.Head
	for ;x.Next!=nil;{
		if x.Next.Key==key{
			x.Next=x.Next.Next
			return
		}
		x=x.Next
	}
}

func (L* List) Display(){
	x:=L.Head
	for ;x!=nil;{
		fmt.Print(x.Key)
		fmt.Print("->")
		x=x.Next
	}
	fmt.Println()
}
