package greedy

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var i, j int
	mySort(g)
	mySort(s)

	// fmt.Println(g)
	// fmt.Println(s)

	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}

		j++
	}

	return i
}

func mySort(nums []int) {
	sort.Ints(nums)
	// bbsort(nums)
}

func bbsort(nums []int) []int {
	var i, j int
	for i = 0; i < len(nums)-1; i++ {
		for j = 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums
}
