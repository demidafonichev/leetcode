func minOperations(nums []int, x int) int {
	if x == 0 {
		return 0
	}
	if len(nums) == 0 || (nums[0] > x && nums[len(nums)-1] > x) {
		return -1
	}
	taken_left, taken_right := nums[1:], nums[0:len(nums)-1]
	x_left, x_right := x-nums[0], x-nums[len(nums)-1]

	from_left, from_right := minOperations(taken_left, x_left), minOperations(taken_right, x_right)
	if from_left == -1 && from_right == -1 {
		return -1
	} else if from_left == -1 {
		return from_right + 1
	} else if from_right == -1 {
		return from_left + 1
	} else {
		if from_left < from_right {
			return from_left + 1
		} else {
			return from_right + 1
		}
	}
}
