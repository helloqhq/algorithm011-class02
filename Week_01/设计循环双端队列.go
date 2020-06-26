package Week_01

type MyCircularDeque struct {
	Head  *Node //头节点
	Tail  *Node //尾结点
	Len   int   //队列长度
	Count int   //节点数量
}

type Node struct {
	Next *Node
	Pre  *Node
	Val  int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	head := Node{
		Next: nil,
		Pre:  nil,
		Val:  -1,
	}
	tail := Node{
		Next: nil,
		Pre:  &head,
		Val:  -1,
	}
	head.Next = &tail

	return MyCircularDeque{
		Head:  &head,
		Tail:  &tail,
		Len:   k,
		Count: 0,
	}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := Node{
		Next: this.Head,
		Pre:  nil,
		Val:  value,
	}
	this.Head.Pre = &newNode
	this.Head = &newNode
	this.Count++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	newNode := Node{
		Next: nil,
		Pre:  this.Tail,
		Val:  value,
	}
	this.Tail.Next = &newNode
	this.Tail = &newNode
	this.Count++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.Head = this.Head.Next
	this.Count--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.Tail = this.Tail.Pre
	this.Count--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	return this.Head.Val
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	return this.Tail.Val
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.Count == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.Count == this.Len {
		return true
	}
	return false
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */