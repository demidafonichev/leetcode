func tupleSameProduct(nums []int) int {
	freq, n, res := map[int]int{}, len(nums), 0

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			prod := nums[i] * nums[j]
			res += 8 * freq[prod]
			freq[prod]++
		}
	}

	return res
}
