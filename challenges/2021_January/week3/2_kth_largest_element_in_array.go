import "sort"

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return -1
	}

	sort.Ints(nums)
	return nums[len(nums)-k]
}
