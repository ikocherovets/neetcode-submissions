func longestCommonPrefix(strs []string) string {
    res := ""

	for i := 0; i < len(strs[0]); i++ {
		for _, str := range strs {
			if i == len(str) || str[i] != strs[0][i] {
				return res
			}
		}
		res += string(strs[0][i])
	}

	return res
}
