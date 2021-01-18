func maxOperations(nums []int, k int) int {
	m := map[int]int{}

	for _, x := range nums {
		m[x]++
	}

	res := 0
	for x := range m {
		if x <= k/2 {
			if x == k/2 && k%2 == 0 {
				res += m[x] / 2
			} else if m[x] < m[k-x] {
				res += m[x]
			} else {
				res += m[k-x]
			}
		}
	}

	return res
}
