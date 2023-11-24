package main

func removeElements(head *ListNode, val int) *ListNode {
	if head != nil && head.Val == val {
		head = head.Next
	}
	current := head
	previous := head
	for current != nil {
		if current.Val == val {
			previous.Next = current.Next
		} else {
			previous = current
		}
		current = current.Next
	}
	if head != nil && head.Val == val {
		head = head.Next
	}
	return head
}
