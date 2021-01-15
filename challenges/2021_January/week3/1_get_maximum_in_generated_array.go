func getMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 2 || n == 1 {
		return 1
	}
	max := 0
	nums := make([]int, n+1)
	nums[1] = 1
	for j := 2; j < n+1; j++ {
		i := j / 2
		if j%2 == 0 {
			nums[j] = nums[i]
		} else {
			nums[j] = nums[i] + nums[i+1]
		}
		if nums[j] > max {
			max = nums[j]
		}
	}
	return max
}
