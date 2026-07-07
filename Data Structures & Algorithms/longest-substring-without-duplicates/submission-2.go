func lengthOfLongestSubstring(s string) int {
    store := make(map[rune]int)
    windowStartIndex := 0
    maxWindowSize := 0

    for currentIndex, char := range s {
        if prevIndex, ok := store[char]; ok && prevIndex >= windowStartIndex {
            windowStartIndex = prevIndex + 1
        }

        store[char] = currentIndex

        currentWindowSize := currentIndex - windowStartIndex + 1
        if currentWindowSize > maxWindowSize {
            maxWindowSize = currentWindowSize
        }
    }

    return maxWindowSize
}