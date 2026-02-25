package piscine

func Map(f func(int) bool, a []int) []bool {
	if a == nil {
		return nil
	}
	res := make([]bool, len(a))
	for i, v := range a {
		res[i] = f(v)
	}
	return res
}
