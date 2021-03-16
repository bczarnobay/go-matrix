package services

func Invert(m [][]string) {
	for r := 0; r < len(m); r++ {
		for c := 0; c < r; c++ {
			var tmp = m[r][c]
			m[r][c] = m[c][r]
			m[c][r] = tmp
		}
	}
}

func Flatten(m [][]string) []string {
	var flat []string
	for r := 0; r < len(m); r++ {
		flat = append(flat, m[r]...)
	}
	return flat
}
