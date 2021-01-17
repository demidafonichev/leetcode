func countGoodRectangles(rectangles [][]int) int {
	count := 0
	maxLen := 0
	for _, rect := range rectangles {
		minSide := rect[0]
		if rect[1] < rect[0] {
			minSide = rect[1]
		}
		if minSide > maxLen {
			maxLen = minSide
			count = 1
		} else if minSide == maxLen {
			count++
		}
	}
	return count
}
