func isPalindrome(head *ListNode) bool {
	n := 0
	for curr := head; curr != nil; curr = curr.Next {
		n++
	}

	middleIndex := n / 2
	if n%2 == 1 {
		middleIndex++
	}

	head1, step := head, 1
	var prev, head2 *ListNode
	for head2 = head; (step - 1) < middleIndex; head2 = head2.Next {
		step++
		prev = head2
	}

	prev = nil
	for curr := head2; curr != nil; {
		next := curr.Next
		curr.Next = prev
		prev, curr = curr, next
	}
	head2 = prev

	for curr1, curr2 := head1, head2; curr1 != nil && curr2 != nil; {
		if curr1.Val != curr2.Val {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}
	return true
}
