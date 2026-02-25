package piscine

func Abort(a, b, c, d, e int) int {
	x := []int{a, b, c, d, e}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4-i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x[2]
}
