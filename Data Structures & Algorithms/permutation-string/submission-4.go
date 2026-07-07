func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	var need, window [26]int // [0,0,0,0,0,0,0,0,0,0,0,...0]

	for i := 0; i < len(s1); i++ {
		need[s1[i] - 'a']++
		window[s2[i] - 'a']++
	}
	// може бути так що відразу все буде, наприклад abs i sab -> тоді паспорта будуть однакові
	if need == window {
		return true
	}

	// якщо ні - треба починати рухати вікно
	for j := len(s1); j < len(s2); j++ { // [1,0,0,0,2] 
		window[s2[j - len(s1)] - 'a']--
		window[s2[j] - 'a']++

		if window == need {
			return true
		}
	}

	return false

}
