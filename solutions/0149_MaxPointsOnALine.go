package solutions

func createFraction(numerator int, denominator int) [2]int {
	if numerator == 0 {
		return [2]int{0, 1}
	}

	factor := 1
	if numerator < 0 {
		numerator = -numerator
		factor *= -1
	}
	if denominator < 0 {
		denominator = -denominator
		factor *= -1
	}

	if denominator == 0 {
		return [2]int{factor, 0}
	}

	n1, n2 := numerator, denominator
	if numerator > denominator {
		n1, n2 = n2, n1
	}

	for n2%n1 != 0 {
		n1, n2 = n2%n1, n1
	}
	return [2]int{factor * numerator / n1, denominator / n1}
}

func MaxPoints(points [][]int) int {
	if len(points) < 2 {
		return len(points)
	}
	m := map[[4]int]map[int]struct{}{}
	for xIdx := 0; xIdx < len(points); xIdx++ {
		xPoint := points[xIdx]
		for yIdx := xIdx + 1; yIdx < len(points); yIdx++ {
			yPoint := points[yIdx]
			slope := createFraction(yPoint[1]-xPoint[1], yPoint[0]-xPoint[0])
			intercept := [2]int{xPoint[0], 0}
			if xPoint[0]-yPoint[0] != 0 {
				intercept = createFraction(xPoint[0]*yPoint[1]-yPoint[0]*xPoint[1], xPoint[0]-yPoint[0])
			}
			key := [4]int{slope[0], slope[1], intercept[0], intercept[1]}
			if m[key] == nil {
				m[key] = map[int]struct{}{}
			}
			m[key][xIdx] = struct{}{}
			m[key][yIdx] = struct{}{}
		}
	}

	res := 0
	for _, value := range m {
		res = max(res, len(value))
	}
	return res
}
