package leetcode

import "sort"

func ContainsDuplicate(num []int) bool {
	sort.Ints(num)
	for i := 1; i < len(num); i++ {
		if num[i] == num[i-1] {
			return true
		}
	}
	return false

}
func ContainsDuplicate1(num []int) bool {
	set := map[int]struct{}{}
	for _, v := range num {
		if _, has := set[v]; has {
			return true
		}
		set[v] = struct{}{}
	}
	return false
}
