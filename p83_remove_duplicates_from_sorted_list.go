package main

func deleteDuplicates(head *ListNode) *ListNode {
	h := head
	if h == nil {
		return h
	}
	for h.Next != nil {
		if h.Next.Val == h.Val {
			h.Next = h.Next.Next
		} else {
			h = h.Next
		}
	}

	return head
}
