package test

import (
	"fmt"
	"golangLeetcode/leetcode"
)

func TestTwoSum() {
	a := [...]int{1, 3, 6, 7}

	b := leetcode.TwoSum(a[:], 13)
	c := leetcode.TwoSum1(a[:], 13)
	fmt.Println(b, c)
}

func ContainsDuplicateTest() {

	d := leetcode.ContainsDuplicate([]int{1, 2, 3, 1})
	e := leetcode.ContainsDuplicate1([]int{1, 2, 3, 1})

	fmt.Println(d, e)
}
func MaxSubArrayTest() {
	f := leetcode.MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})

	fmt.Println(f)
}
func PlusOneTest() {
	g := leetcode.PlusOne([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9, 9})
	fmt.Println(g)
}
