package solutions

func FindOrder(num int, prerequisites [][]int) []int {
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

	idx := 0
	for idx != len(queue) {
		id := queue[idx]
		idx++
		for _, next := range graph[id] {
			degrees[next]--
			if degrees[next] == 0 {
				queue = append(queue, next)
				remain--
			}
		}
	}

	if remain == 0 {
		return queue
	} else {
		return []int{}
	}
}
