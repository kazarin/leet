package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type MyLinkedList struct {
	root *ListNode
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) int {
	root := this.root
	for i := 0; i < index; i++ {
		root = root.Next
		if root == nil {
			return -1
		}
	}
	if root == nil {
		return -1
	}
	return root.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	tmp := this.root
	this.root = &ListNode{Val: val, Next: tmp}
}

func (this *MyLinkedList) AddAtTail(val int) {
	root := this.root
	if root == nil {
		this.AddAtHead(val)
		return
	}
	for root.Next != nil {
		root = root.Next

	}
	tmp := &ListNode{Val: val}
	root.Next = tmp

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	root := this.root
	if index == 0 {
		this.AddAtHead(val)
		return

	}
	if root == nil {
		return
	}
	for i := 0; i < index-1; i++ {
		root = root.Next
		if root == nil {
			return
		}

	}
	if root == nil {
		return
	}
	tmp := &ListNode{Val: val, Next: root.Next}
	root.Next = tmp

}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	root := this.root
	if index == 0 {
		this.root = root.Next
		return
	}
	for i := 0; i < index-1; i++ {
		root = root.Next

	}
	if root == nil {
		return
	}
	if root.Next != nil {
		root.Next = root.Next.Next
	}

}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
