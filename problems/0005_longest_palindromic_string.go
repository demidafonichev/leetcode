func longestPalindrome(s string) string {
	n := len(s)
	for l := n; l >= 0; l-- {
		for i := 0; i <= n-l; i++ {
			mi, pal := i+l/2, true
			for j := i; j <= mi; j++ {
				if s[j] != s[(i+l-1)-(j-i)] {
					pal = false
					break
				}
			}
			if pal {
				return s[i : i+l]
			}
		}
	}

	return ""
}
