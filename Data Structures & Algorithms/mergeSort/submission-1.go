// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }
func merge(arr []Pair, l, r, m int) {
	// arr[l..m]   — ліва половина (m належить лівій половині!)
	// arr[m+1..r] — права половина (починається саме з m+1)
	//
	//   індекси:   l  l+1 ... m | m+1 ... r
	//              |--- ліва ---|--- права ---|

	leftSubArr := make([]Pair, m-l+1) // довжина лівої половини: від l до m включно
	rightSubArr := make([]Pair, r-m)  // довжина правої половини: від m+1 до r включно

	// Копіюємо ліву половину: arr[l] відповідає leftSubArr[0],
	// тому зсув простий — просто +l
	for i := 0; i < len(leftSubArr); i++ {
		leftSubArr[i] = arr[i+l]
	}

	// Копіюємо праву половину: перший елемент правої половини — це arr[m+1],
	// а не arr[m] (бо arr[m] уже належить лівій половині і скопійований вище).
	// Тому при j=0 треба отримати arr[m+1], звідси зсув m+1:
	//   j=0 → arr[m+1] (перший елемент правої частини)
	//   j=1 → arr[m+2]
	//   ...
	//   j=r-m-1 → arr[r] (останній елемент, бо довжина = r-m)
	for j := 0; j < len(rightSubArr); j++ {
		rightSubArr[j] = arr[j+m+1]
	}

	i := 0
	j := 0
	k := l // k — позиція запису назад в оригінальний arr, починаємо з l (не з 0!)

	// Стандартне злиття двох відсортованих підмасивів
	for i < len(leftSubArr) && j < len(rightSubArr) {
		if leftSubArr[i].Key <= rightSubArr[j].Key {
			// <= (а не <) забезпечує стабільність сортування:
			// при рівних ключах елемент з лівої половини (тобто той,
			// що йшов раніше в оригінальному масиві) береться першим
			arr[k] = leftSubArr[i]
			i++
		} else {
			arr[k] = rightSubArr[j]
			j++
		}
		k++
	}

	// Дописуємо залишок лівої половини, якщо права закінчилась раніше
	for i < len(leftSubArr) {
		arr[k] = leftSubArr[i]
		i++
		k++
	}

	// Дописуємо залишок правої половини, якщо ліва закінчилась раніше
	for j < len(rightSubArr) {
		arr[k] = rightSubArr[j]
		j++
		k++
	}
}

func mergeSort(pairs []Pair) []Pair {
	var mergeSortMain func(arr []Pair, l, r int)
	mergeSortMain = func(arr []Pair, l, r int) {
		if r-l+1 <= 1 {
			// підмасив з 0 або 1 елементів — уже "відсортований"
			return
		}
		m := l + (r-l)/2 // середина; такий запис уникає переповнення int
		                 // порівняно з (l+r)/2

		mergeSortMain(arr, l, m)   // сортуємо ліву частину: [l..m]
		mergeSortMain(arr, m+1, r) // сортуємо праву частину: [m+1..r]
		merge(arr, l, r, m)        // зливаємо обидві відсортовані частини
	}
	mergeSortMain(pairs, 0, len(pairs)-1)
	return pairs
}