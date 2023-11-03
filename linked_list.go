package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) PrettyPrint() {
	head := l

	for head.Next != nil {
		fmt.Printf("%d -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("%d", head.Val)
	fmt.Println()
}

func (l *ListNode) ToArray() []int {
	head := l
	var res []int

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func IntsToListNodes(nums []int) *ListNode {
	h := &ListNode{Val: nums[0]}
	head := h
	for i := 1; i < len(nums); i++ {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &ListNode{Val: nums[i]}

	}
	return h
}

func IntsToListNodesWithCycle(nums []int, pos int) *ListNode {
	head := IntsToListNodes(nums)
	if pos == -1 {
		return head
	}

	cycleTo := &ListNode{}

	i := 0
	for head.Next != nil {
		if i == pos {
			cycleTo = head
		}
		head = head.Next
		i++
	}
	head.Next = cycleTo
	return head
}

func (list1 *ListNode) SameAs(list2 *ListNode) bool {
	l1 := list1
	l2 := list2
	for l1.Next != nil && l2.Next != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next

	}
	return true

}
