package greedy

import (
	"sort"
)

func findContentChildren(g []int, s []int) int {
	var i, child, lastCookie, contentCount, cookieLen int
	cookieLen = len(s)
	mySort(g)
	mySort(s)

	// fmt.Println(g)
	// fmt.Println(s)

	lastCookie = -1

	contentCount = 0

	for _, child = range g {
		// find suitable cookie
		for i = lastCookie + 1; i < cookieLen; i++ {
			if child > s[i] {

			} else {
				contentCount++
				lastCookie = i
				break
			}
		}
	}

	return contentCount
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
