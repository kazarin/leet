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
