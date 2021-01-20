func lengthOfLongestSubstring(s string) int {
	max, l := 0, 0
	used := make([]bool, 128)
	posi := make([]int, 128)
	for i := 0; i < len(s); i++ {
		char := s[i]
		if used[char] {
			if l > max {
				max = l
			}
			l = 0
			i = posi[char]
			used = make([]bool, 128)
			posi = make([]int, 128)
			continue
		}
		l++
		used[char] = true
		posi[char] = i
	}
	if l > max {
		max = l
	}
	return max
}
