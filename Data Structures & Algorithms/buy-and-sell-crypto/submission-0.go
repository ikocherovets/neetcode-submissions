func maxProfit(prices []int) int {
    res := 0

    for i := 0; i < len(prices); i++ {
        buy := prices[i]

        for j := i + 1; j < len(prices); j++ {
            sell := prices[j]
            res = int(math.Max(float64(res), float64(sell-buy)))
        }
    }

    return res
}