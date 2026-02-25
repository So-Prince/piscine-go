package piscine

func Capitalize(s string) string {
	b := []byte(s)
	capNext := true

	for i := 0; i < len(b); i++ {
		c := b[i]

		isLetter := (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
		isDigit := (c >= '0' && c <= '9')

		if isLetter || isDigit {
			if capNext && isLetter {
				if c >= 'a' && c <= 'z' {
					b[i] = c - ('a' - 'A')
				}
				capNext = false
			} else if !capNext && isLetter {
				if c >= 'A' && c <= 'Z' {
					b[i] = c + ('a' - 'A')
				}
			} else {
				// digit, do nothing
			}
			capNext = false
		} else {
			capNext = true
		}
	}
	return string(b)
}
