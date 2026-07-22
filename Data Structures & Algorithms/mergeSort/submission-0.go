// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }
func merge(arr []Pair, l, r, m int) {
	leftSubArr := make([]Pair, m-l+1)
	rightSubArr := make([]Pair, r-m)
	for i := 0; i < len(leftSubArr); i++ {
		leftSubArr[i] = arr[i+l]
	}
	for j := 0; j < len(rightSubArr); j++ {
		rightSubArr[j] = arr[j+m+1]
	}
	i := 0
	j := 0
	k := l
	for i < len(leftSubArr) && j < len(rightSubArr) {
		if leftSubArr[i].Key <= rightSubArr[j].Key {
			arr[k] = leftSubArr[i]
			i++
		} else {
			arr[k] = rightSubArr[j]
			j++
		}
		k++
	}
	for i < len(leftSubArr) {
		arr[k] = leftSubArr[i]
		i++
		k++
	}
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
			return
		}
		m := l + (r-l)/2 // to avoid int overflow
		mergeSortMain(arr, l, m)
		mergeSortMain(arr, m+1, r)
		merge(arr, l, r, m)
	}
	mergeSortMain(pairs, 0, len(pairs)-1)
	return pairs
}