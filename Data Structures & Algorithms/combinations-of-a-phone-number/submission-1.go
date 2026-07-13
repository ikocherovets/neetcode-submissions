func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	result := []string{}

	// Контракт backtrack(i, curString):
	//   i         — індекс цифри, яку зараз обробляємо
	//   curString — рядок, зібраний з букв попередніх цифр
	//   base case: len(curString) == len(digits) — комбінація готова, зберігаємо
	//   recursive case: перебираємо всі букви цифри digits[i], для кожної йдемо на i+1

	digitMap := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(i int, curString string)
	backtrack = func(i int, curString string) {
		// base case: набрали стільки ж букв, скільки цифр — комбінація готова
		if len(curString) == len(digits) {
			result = append(result, curString)
			return
		}

		// recursive case: беремо букви поточної цифри digits[i]
		// і для кожної окремо рекурсивно добудовуємо рядок далі
		for _, c := range digitMap[digits[i]] {
			backtrack(i+1, curString+string(c))
		}
	}

	backtrack(0, "")

	return result
}