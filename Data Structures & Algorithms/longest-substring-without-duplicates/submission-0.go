func lengthOfLongestSubstring(s string) int {
    charLastIndex := make(map[rune]int) // тут ми зберігаємо значення останнього індекса. 
    windowStart := 0
    longestWindow := 0

    for windowEnd, char := range s {
        if prevIndex, ok := charLastIndex[char]; ok && prevIndex >= windowStart {
            windowStart = prevIndex + 1
        }

        charLastIndex[char] = windowEnd

        currentWindowSize := windowEnd - windowStart + 1
        if currentWindowSize > longestWindow {
            longestWindow = currentWindowSize
        }
    }

    return longestWindow
} 