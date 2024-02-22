package solutions

type MinStack struct {
	size     int
	stack    []int
	valStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{size: 0, stack: []int{}, valStack: []int{}}
}

func (s *MinStack) Push(val int) {
	if s.size == 0 {
		s.stack = append(s.stack, val)
	} else {
		s.stack = append(s.stack, min(val, s.stack[s.size-1]))
	}
	s.valStack = append(s.valStack, val)
	s.size++
}

func (s *MinStack) Pop() {
	if s.size != 0 {
		s.size--
		s.stack = s.stack[:s.size]
		s.valStack = s.valStack[:s.size]
	}
}

func (s *MinStack) Top() int {
	return s.valStack[s.size-1]
}

func (s *MinStack) GetMin() int {
	return s.stack[s.size-1]
}
