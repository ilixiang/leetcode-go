package solutions

func CanFinish(num int, prerequisites [][]int) bool {
	graph := make([][]int, num)
	degrees := make([]int, num)
	for _, prerequisite := range prerequisites {
		to, from := prerequisite[0], prerequisite[1]
		graph[from] = append(graph[from], to)
		degrees[to]++
	}

	remain := 0
	queue := make([]int, 0, num)
	for id, degree := range degrees {
		if degree == 0 {
			queue = append(queue, id)
		} else {
			remain++
		}
	}

	for len(queue) != 0 {
		id := queue[0]
		queue = queue[1:]
		for _, next := range graph[id] {
			degrees[next]--
			if degrees[next] == 0 {
				queue = append(queue, next)
				remain--
			}
		}
	}

	return remain == 0
}
