package utils

func Intersect(a, b []int) (c []int) {
	m := make(map[int]int)

	for _, v := range a {
		m[v] = v
	}

	for _, v := range b {
		if _, ok := m[v]; ok {
			c = append(c, v)
		}
	}

	return
}
