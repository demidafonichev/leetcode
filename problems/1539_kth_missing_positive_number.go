func findKthPositive(arr []int, k int) int {
	kth, idx := 0, 0
	num := 1
	for {
		if idx < len(arr) && num == arr[idx] {
			idx++
		} else {
			kth++
			if kth == k {
				return num
			}
		}
		num++
	}
}

func findKthPositiveAlternative(arr []int, k int) int {
	kth, idx := 0, 0
	for i := 1; i <= 2000; i++ {
		if idx < len(arr) && i == arr[idx] {
			idx++
		} else {
			kth++
			if kth == k {
				return i
			}
		}
	}
	return 0
}
