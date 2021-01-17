func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	n, xcopy := 0, x
	for x != 0 {
		n = n*10 + x%10
		x /= 10
	}
	return n == xcopy
}
