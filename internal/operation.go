package internal

func Intersection(a []int, b []int) []int {
	minLen := min(len(a), len(b))
	r := make([]int, 0, minLen)

	var i, j int
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			i++
		} else if a[i] > b[j] {
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}
	return r
}

func Union(a []int, b []int) []int {
	r := make([]int, 0, len(a)+len(b))

	var i, j int

	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			r = append(r, a[i])
			i++
		} else if a[i] > b[j] {
			r = append(r, b[j])
			j++
		} else {
			r = append(r, a[i])
			i++
			j++
		}
	}

	for i < len(a) {
		r = append(r, a[i])
		i++
	}

	for j < len(b) {
		r = append(r, b[j])
		j++
	}

	return r
}
