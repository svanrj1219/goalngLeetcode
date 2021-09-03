package leetcode

func PlusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] = digits[i] + 1
		if digits[i] != 10 {
			return digits
		}
		digits[i] = digits[i] % 10
	}
	return append([]int{1}, digits...)
}
