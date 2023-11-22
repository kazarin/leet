package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	h := head
	l := h
	r := h

	for i := 0; i < n; i++ {
		r = r.Next
	}
	if r == nil {
		return h.Next
	}

	for r != nil && r.Next != nil {
		r = r.Next
		l = l.Next
	}
	l.Next = l.Next.Next

	return h
}
