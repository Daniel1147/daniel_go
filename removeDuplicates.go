package main

import "fmt"

func removeDuplicates(nums []int) int {
	var last int
	num := 0
	trueIndex := 0
	firstFlag := true
	for _, value := range nums {
		if (firstFlag == true) {
            firstFlag = false
			last = value
			num++
			trueIndex++
		} else if (last != value) {
			nums[trueIndex] = value
            last = value
			num++
			trueIndex++
		}
	}

	nums = nums[:trueIndex]

	return num
}

func main() {
	input := []int{1, 2, 2}

	fmt.Println(removeDuplicates(input))
    fmt.Println(input)

	return
}

