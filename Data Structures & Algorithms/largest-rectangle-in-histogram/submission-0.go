func largestRectangleArea(heights []int) int {
    type Bar struct {
        index  int
        height int
    }

    stack := []Bar{}
    maxArea := 0

    // sentinel 0 at the end forces all remaining bars to be popped
    for i, h := range append(heights, 0) {
        // tracks where the current bar "could have started"
        startIndex := i

        for len(stack) > 0 && stack[len(stack)-1].height >= h {
            top := stack[len(stack)-1]
            stack = stack[:len(stack)-1]

            area := top.height * (i - top.index)
            if area > maxArea {
                maxArea = area
            }

            // this bar extends back to where the popped bar started
            startIndex = top.index
        }

        stack = append(stack, Bar{index: startIndex, height: h})
    }

    return maxArea
}