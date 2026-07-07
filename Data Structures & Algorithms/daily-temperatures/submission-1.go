func dailyTemperatures(temperatures []int) []int {
    result := make([]int, len(temperatures))
    // monotonic decreasing stack — stores indices of days
    // with unresolved "next warmer day"
    pendingIndices := []int{}

    for today, temp := range temperatures {
        // current temp is warmer than days waiting on the stack
        for len(pendingIndices) > 0 {
            prevDay := pendingIndices[len(pendingIndices)-1]
            if temp <= temperatures[prevDay] {
                break
            }
            pendingIndices = pendingIndices[:len(pendingIndices)-1]
            result[prevDay] = today - prevDay
        }
        pendingIndices = append(pendingIndices, today)
    }

    return result
}