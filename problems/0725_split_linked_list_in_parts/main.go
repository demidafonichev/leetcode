func splitListToParts(root *ListNode, k int) (res []*ListNode) {
	n := 0
	for curr := root; curr != nil; curr = curr.Next {
		n++
	}

	chunkSize := n / k
	chunksRemaining := k
	inResult := 0
	inResultMax := 0
	if n > k {
		inResultMax = n
	} else {
		inResultMax = k
	}
	for curr := root; inResult < inResultMax; {
		currentChunkSize := chunkSize
		if n != 0 && (n-inResult)%(chunksRemaining) != 0 {
			currentChunkSize++
		}

		res = append(res, curr)
		inResult++
		for i := 0; i < currentChunkSize-1; i++ {
			curr = curr.Next
			inResult++
		}
		if curr != nil {
			next := curr.Next
			curr.Next = nil
			curr = next
		}
		chunksRemaining--
	}

	return res
}
