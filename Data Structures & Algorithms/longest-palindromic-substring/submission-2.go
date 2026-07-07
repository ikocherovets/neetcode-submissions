func longestPalindrome(s string) string {
    if len(s) < 1 {
        return ""
    }

    // повертає межі знайденого паліндрома [left, right]
    expand := func(left, right int) (int, int) {
        for left >= 0 && right < len(s) && s[left] == s[right] {
            left--
            right++
        }
        return left + 1, right - 1
    }

    start, end := 0, 0
    for i := 0; i < len(s); i++ {
        l1, r1 := expand(i, i)   // непарний центр
        l2, r2 := expand(i, i+1) // парний центр

        if r1-l1 > end-start {
            start, end = l1, r1
        }
        if r2-l2 > end-start {
            start, end = l2, r2
        }
    }

    return s[start : end+1]
}