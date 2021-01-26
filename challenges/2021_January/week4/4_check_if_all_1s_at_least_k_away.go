func kLengthApart(nums []int, k int) bool {
	n, prev := len(nums), -k-1

	for i := 0; i < n; i++ {
		if nums[i] == 1 {
			if i-prev < k+1 {
				return false
			}
			prev = i
		}
	}
	return true
}
