func mostCompetitive(nums []int, k int) []int {
	n, res, idx := len(nums), make([]int, k), 0
	res[0] = nums[0]

	for i := 1; i < n; i++ {
		putFront := false
		for j, pos := idx, idx; j >= 0 && (n-i) >= (k-pos); j-- {
			if nums[i] < res[j] {
				res[j] = nums[i]
				idx = j
				pos--
				putFront = true
			} else {
				break
			}
		}
		if !putFront && idx < k-1 {
			idx++
			res[idx] = nums[i]
		}
	}
	return res
}
