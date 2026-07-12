func minEatingSpeed(piles []int, h int) int {
    l, r := 1, 0

    // шукаємо максимальну купу
    for _, p := range piles {
        if p > r {
            r = p
        }
    }

    res := r

    // binary search по швидкості поїдання
    // діапазон: від 1 банана/год до найбільшої купи
    for l <= r {

        // пробуємо середню швидкість
        k := (l + r) / 2

        totalTime := 0

        // рахуємо, скільки годин займе поїдання
        // при швидкості k
        for _, p := range piles {
            totalTime += int(math.Ceil(float64(p) / float64(k)))
        }

        if totalTime <= h {
            // ця швидкість працює,
            // але шукаємо ще меншу
            res = k
            r = k - 1
        } else {
            // занадто повільно,
            // треба збільшити швидкість
            l = k + 1
        }
    }

    return res
}