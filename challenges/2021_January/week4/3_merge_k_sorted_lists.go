func mergeKLists(lists []*ListNode) *ListNode {
	sub := &ListNode{}
	prev := sub

	for {
		minVal, minValIdx := 10001, -1
		for i, head := range lists {
			if head != nil && head.Val < minVal {
				minVal = head.Val
				minValIdx = i
			}
		}
		if minValIdx == -1 {
			break
		}

		prev.Next = lists[minValIdx]
		prev = lists[minValIdx]
		lists[minValIdx] = lists[minValIdx].Next
	}

	return sub.Next
}