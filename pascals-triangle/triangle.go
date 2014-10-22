package pascal

var triangles = [][]int{{1}}

func init() {
	for j := 0; j < 30; j++ {
		triangles = append(triangles, next(triangles[len(triangles)-1]))
	}
}

func next(line []int) []int {
	res := make([]int, 0, len(line)+1)
	res = append(res, 1)
	for j := 1; j < len(line); j++ {
		res = append(res, line[j]+line[j-1])
	}
	res = append(res, 1)
	return res
}

func Triangle(n int) [][]int {
	return triangles[:n]
}
