package piscine

func ReverseMenuIndex(menu []string) []string {
	length := 0
	for range menu {
		length++
	}

	reversedMenu := make([]string, length)

	i := 0
	for j := length - 1; j >= 0; j-- {
		reversedMenu[i] = menu[j]
		i++
	}

	return reversedMenu
}
