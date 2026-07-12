func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	var bfs func([]int, int) int
	bfs = func(coins []int, amount int) int {
		queue := []int{0}
		seen := make([]bool, amount+1)
		seen[0] = true
		steps := 0

		for len(queue) > 0 {
			steps++
			size := len(queue)

			for i := 0; i < size; i++ {
				cur := queue[0]
				queue = queue[1:]

				for _, coin := range coins {
					next := cur + coin

					if next == amount {
						return steps
					}
					if next > amount || seen[next] {
						continue
					}

					seen[next] = true
					queue = append(queue, next)
				}
			}
		}

		return -1
	}

	return bfs(coins, amount)
}