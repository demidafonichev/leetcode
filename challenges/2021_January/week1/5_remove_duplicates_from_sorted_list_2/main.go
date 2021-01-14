func deleteDuplicates(head *ListNode) *ListNode {
	prehead := &ListNode{-101, head}
	l := prehead

	for head != nil {
		if head.Next != nil && head.Val == head.Next.Val {
			for head.Next != nil && head.Val == head.Next.Val {
				head = head.Next
			}
			l.Next = head.Next
		} else {
			l = l.Next
		}
		head = head.Next
	}
	return prehead.Next
}
