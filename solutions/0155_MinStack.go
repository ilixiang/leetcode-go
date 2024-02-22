package solutions

type MinStack struct {
	size     int
	stack    []int
	valStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{size: 0, stack: []int{}, valStack: []int{}}
}

func (this *MinStack) Push(val int) {
	if this.size == 0 {
		this.stack = append(this.stack, val)
	} else {
		this.stack = append(this.stack, min(val, this.stack[this.size-1]))
	}
	this.valStack = append(this.valStack, val)
	this.size++
}

func (this *MinStack) Pop() {
	if this.size != 0 {
		this.size--
		this.stack = this.stack[:this.size]
		this.valStack = this.valStack[:this.size]
	}
}

func (this *MinStack) Top() int {
	return this.valStack[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.size-1]
}
