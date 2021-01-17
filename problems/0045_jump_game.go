func jump(nums []int) int {
	n := len(nums)
	minJumps := make([]int, n)
	for i := 0; i < n; i++ {
		minJumps[i] = 100000
	}
	minJumps[0] = 0

	for i := 0; i < n; i++ {
		for j := 1; j <= nums[i] && (i+j) < n; j++ {
			if minJumps[i]+1 < minJumps[i+j] {
				minJumps[i+j] = minJumps[i] + 1
			}
		}
	}
	return minJumps[len(nums)-1]
}
