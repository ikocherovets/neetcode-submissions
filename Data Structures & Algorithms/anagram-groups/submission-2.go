func groupAnagrams(strs []string) [][]string {
    var result [][]string
    store := make(map[string][]string)

    for _, s := range strs {
        sortedS := sortString(s)
        store[sortedS] = append(store[sortedS], s)
    }

    for _, group := range store {
        result = append(result, group)
    }

    return result 
}


func sortString(word string) string {
    chars := []rune(word)

    sort.Slice(chars, func(i, j int) bool {
        return chars[i] < chars[j]
    })

    return string(chars)
}
