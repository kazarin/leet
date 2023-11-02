package main

func hasCycle(head *ListNode) bool {
	visited := map[ListNode]int{}
	for head != nil && head.Next != nil {
		if _, ok := visited[*head]; ok {
			return true
		}
		visited[*head]++
		head = head.Next
	}
	return false
}
