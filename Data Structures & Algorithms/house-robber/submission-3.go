func rob(nums []int) int {
	type Cache map[int]int
	cache := make(Cache)
	// контракт - dfs(i) = максимальна сума, яку можна отримати, починаючи з будинку i.
	// ця сума це сума або сума від n + 2 + цей будинок, або вже від наступного
    var dfs func (i int, cache Cache) int
	dfs = func(i int, cache Cache) int {
		// чому >= бо ми перепригнем на +1 індекс і щоб уникнути index out of range 
		if i >= len(nums) {
			return 0
		}

		if val, ok := cache[i]; ok {
			return val
		}

		result := max(dfs(i + 1, cache), nums[i] + dfs(i + 2, cache))
		cache[i] = result
		return result 
	}
	
	return dfs(0, cache)
}
