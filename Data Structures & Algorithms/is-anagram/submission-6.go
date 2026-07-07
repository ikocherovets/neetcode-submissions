func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    type charStore map[rune]int
    store1 := make(charStore)
    store2 := make(charStore)

    for i, char := range s { // чому відразу два лупи - бо довжина слів має бути однакова якщо це анаграми
        store1[char]++
        store2[rune(t[i])]++
    }

    for key, value := range store1 {
        if value != store2[key] {
            return false
        }
    }

    return true
}
