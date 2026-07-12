func searchMatrix(matrix [][]int, target int) bool {
    rows, cols := len(matrix), len(matrix[0])

    top, bottom := 0, rows-1

    // Шукаємо потрібний рядок
    for top <= bottom {
        row := (top + bottom) / 2

        if target > matrix[row][cols-1] {
            top = row + 1
        } else if target < matrix[row][0] {
            bottom = row - 1
        } else {
            break
        }
    }

    if top > bottom {
        return false
    }

    row := (top + bottom) / 2
    left, right := 0, cols-1

    // Бінарний пошук у знайденому рядку
    for left <= right {
        mid := (left + right) / 2

        if matrix[row][mid] == target {
            return true
        } else if matrix[row][mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }

    return false
}