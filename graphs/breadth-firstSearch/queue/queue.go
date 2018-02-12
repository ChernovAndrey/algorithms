package queue

const N =4 //начальный размер
type Queue struct{
	array []interface{}
	tail int
	head int
}

func (Q *Queue) Init(){
	Q.array=make([]interface{},N)
	Q.tail=0
	Q.head=0
}

func (Q *Queue) expand(){
	n:=len(Q.array)
	nArray := make([]interface{},2*n)
	for i:=0;i<n;i++{
		nArray[i]=Q.array[i]
	}
	Q.array=nArray
}
func (Q *Queue) EnQueue(x interface{}){
	if (Q.tail+Q.head) == (len(Q.array)-1) {
		Q.expand()
	}

	n:=len(Q.array)
	Q.array[Q.tail]=x

	if Q.tail==(n-1){
		Q.tail=0
	}else {
		Q.tail++
	}
}
func (Q *Queue) DeQueue() interface{}{
	res:=Q.array[Q.head]
	if Q.tail==Q.head{
		panic("queue is Empty")
	}
	if Q.head==(len(Q.array)-1){
		Q.head=0
	}else {
		Q.head++
	}
	return res
}

func (Q *Queue) IsEmpty() bool{
	return Q.head==Q.tail
}