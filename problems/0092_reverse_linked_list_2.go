func reverseBetween(head *ListNode, m int, n int) *ListNode {
	n = n - m
	prehead := &ListNode{0, head}

	preH, H := prehead, head
	for m > 1 {
		preH = H
		H = H.Next
		m--
	}

	prev, T, postT := preH, H, H.Next
	for n > 0 {
		prev, T = T, postT
		postT = postT.Next
		T.Next = prev
		n--
	}

	preH.Next = T
	H.Next = postT

	return prehead.Next
}
