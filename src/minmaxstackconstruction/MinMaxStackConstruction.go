package minmaxstackconstruction

type MinMaxStack struct {
	// Write your code here.
	top       *StackFrame
	minMaxTop *MinMaxFrame
}

type StackFrame struct {
	value int
	next  *StackFrame
}

type MinMaxFrame struct {
	min  int
	max  int
	next *MinMaxFrame
}

func NewMinMaxStack() *MinMaxStack {
	// Write your code here.
	return &MinMaxStack{}
}

func (stack *MinMaxStack) Peek() int {
	// Write your code here.
	if stack.top != nil {
		return stack.top.value
	}
	return -1
}

func (stack *MinMaxStack) Pop() int {
	// Write your code here.
	value := -1
	if stack.top != nil {
		value = stack.top.value
		stack.top = stack.top.next
		stack.minMaxTop = stack.minMaxTop.next
	}
	return value
}

func (stack *MinMaxStack) Push(number int) {
	// Write your code here.
	if stack.top == nil {
		stack.top = &StackFrame{
			value: number,
			next:  nil,
		}
		stack.minMaxTop = &MinMaxFrame{
			min: number,
			max: number,
		}

	} else {
		temp := stack.top
		stack.top = &StackFrame{
			value: number,
			next:  temp,
		}
		tempMinMax := stack.minMaxTop
		min := number
		max := number
		if tempMinMax.max > max {
			max = tempMinMax.max
		}

		if tempMinMax.min < min {
			min = tempMinMax.min
		}
		stack.minMaxTop = &MinMaxFrame{
			min:  min,
			max:  max,
			next: tempMinMax,
		}

	}
}

func (stack *MinMaxStack) GetMin() int {
	// Write your code here.
	if stack.minMaxTop != nil {
		return stack.minMaxTop.min
	}
	return -1
}

func (stack *MinMaxStack) GetMax() int {
	// Write your code here.
	if stack.minMaxTop != nil {
		return stack.minMaxTop.max
	}
	return -1
}
