package piscine

func PodiumPosition(podium [][]string) [][]string {
	length := 0
	for range podium {
		length++
	}

	i := 0
	j := length - 1
	for i < j {
		temp := podium[i]
		podium[i] = podium[j]
		podium[j] = temp
		i++
		j--
	}

	return podium
}
