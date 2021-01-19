func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	for cur := prehead; l1 != nil || l2 != nil; cur = cur.Next {
		var next ListNode
		if l1 != nil && l2 != nil && l1.Val <= l2.Val || l1 != nil && l2 == nil {
			next = ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			next = ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		cur.Next = &next
	}
	return prehead.Next
}
