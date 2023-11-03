package main

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	for h.Next != nil {
		if h.Next.Val == h.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}

	return head
}
