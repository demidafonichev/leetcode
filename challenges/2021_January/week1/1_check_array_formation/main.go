func canFormArray(arr []int, pieces [][]int) bool {
	i, j := 0, 0

	for i < len(arr) {
		pFound := false
		for j < len(pieces) {
			ii, v, k, p := i, arr[i], 0, pieces[j]
			for k < len(p) {
				if v == p[k] {
					ii++
					if ii == len(arr) {
						return true
					}
					v = arr[ii]
					k++
					if k == len(p) {
						pFound = true
						break
					}
				} else {
					break
				}
			}
			if pFound {
				pieces[j], pieces[len(pieces)-1] = pieces[len(pieces)-1], pieces[j]
				pieces = pieces[:len(pieces)-1]
				arr = arr[ii:]
				i, j = 0, 0
				break
			}
			j++
		}
		if !pFound {
			return false
		}
	}
	return false
}
