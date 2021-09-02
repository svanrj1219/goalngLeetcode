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
