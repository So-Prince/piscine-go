package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	var result []string

	for _, s := range slice {
		if s != "" {
			result = append(result, s)
		}
	}

	*ptr = result
	return len(result)
}
