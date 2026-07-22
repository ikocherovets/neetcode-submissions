/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left, right := 1, n
	for left <= n {
		midNumber := left + (right - left) / 2 	// to avoit int overflow
		result := guess(midNumber)
	
		if result == -1  {
			right = midNumber - 1
		} else if result == 1 {
			left = midNumber + 1
		} else {
				return midNumber
		}
	}
	return -1 
}
