package piscine

func ListSize(l *List) int {
	counter := 0

	for node := l.Head; node != nil; node = node.Next {
		counter++
	}
	return counter
}
