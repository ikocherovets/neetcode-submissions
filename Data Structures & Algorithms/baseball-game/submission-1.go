func calPoints(operations []string) int {
    stack := []int{}
    for _, op := range operations {
        if op == "+" {
            n := len(stack)
            stack = append(stack, stack[n-1]+stack[n-2])
        } else if op == "D" {
            stack = append(stack, 2*stack[len(stack)-1])
        } else if op == "C" {
            stack = stack[:len(stack)-1]
        } else {
            num, _ := strconv.Atoi(op)
            stack = append(stack, num)
        }
    }
    sum := 0
    for _, v := range stack {
        sum += v
    }
    return sum
}