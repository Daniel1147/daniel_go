package main

import "fmt"

func maxSubArray(nums []int) int {

    largestSum := nums[0]
    lastSum := 0
    for _, num := range nums {
        if (lastSum > 0) {
            lastSum = lastSum + num
        } else {
            lastSum = num
        }

        // fmt.Println(num, lastSum)

        if (largestSum < lastSum) {
            largestSum = lastSum
        }
    }

    return largestSum
}

func main() {
    // input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
    input := []int{-2}

    fmt.Println(maxSubArray(input))

	return
}

