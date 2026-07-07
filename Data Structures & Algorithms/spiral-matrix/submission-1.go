func spiralOrder(matrix [][]int) []int {
    res := []int{}
    left, right := 0, len(matrix[0])
    top, bottom := 0, len(matrix)

    for left < right && top < bottom {
        // → верхній рядок
        for i := left; i < right; i++ {
            res = append(res, matrix[top][i])
        }
        top++

        // ↓ правий стовпець
        for i := top; i < bottom; i++ {
            res = append(res, matrix[i][right-1])
        }
        right--

        // після двох кроків межі могли зійтись — перевіряємо
        if !(left < right && top < bottom) {
            break
        }

        // ← нижній рядок
        for i := right - 1; i >= left; i-- {
            res = append(res, matrix[bottom-1][i])
        }
        bottom--

        // ↑ лівий стовпець
        for i := bottom - 1; i >= top; i-- {
            res = append(res, matrix[i][left])
        }
        left++
    }

    return res
}