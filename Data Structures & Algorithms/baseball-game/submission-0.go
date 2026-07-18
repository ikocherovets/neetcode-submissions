func calPoints(operations []string) int {
	stack := []int{}

	for _, op := range operations {
		number, err := strconv.Atoi(op)
		if err == nil {
			stack = append(stack, number)
		} else {
			switch op {
			case "+":
				a := stack[len(stack) - 1]
				b := stack[len(stack) - 2]
				stack = append(stack, a + b)
			case "D":
				a := stack[len(stack) - 1]
				stack = append(stack, a * 2)
			case "C":
				stack = stack[:len(stack) - 1]
			}
		}
	}

	sum := 0
	for _, num := range stack {
		sum += num
	}

	return sum
}
