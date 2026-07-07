type MinStack struct {
	stack []int
	mins  []int // mins[i] = minimum of stack[0..i]
}

func Constructor() MinStack {
	return MinStack{}
}

func (s *MinStack) Push(val int) {
	s.stack = append(s.stack, val)
	min := val
	if len(s.mins) > 0 && s.mins[len(s.mins)-1] < min {
		min = s.mins[len(s.mins)-1]
	}
	s.mins = append(s.mins, min)
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.mins = s.mins[:len(s.mins)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.mins[len(s.mins)-1]
}