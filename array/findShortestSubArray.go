package array

import "fmt"

func findShortestSubArray(nums []int) int {
    countList := degreeList(nums)
    ans := shortestSubArr(countList, nums)

    return ans
}

func shortestSubArr(countList map[int]int, nums []int) int {
    var numStart, numEnd map[int]int

    maxFrequency := maxElement(countList)

    numStart = make(map[int]int)
    numEnd = make(map[int]int)

    for k, v := range nums {
        if countList[v] == maxFrequency {
            _, ok := numStart[v]
            if ok == false {
                numStart[v] = k
                numEnd[v] = k
            } else {
                numEnd[v] = k
            }
        }
    }

    minLen := -1
    for k, v := range numStart {
        // fmt.Println(numStart[k], numEnd[k])
        length := numEnd[k] - v + 1
        if minLen == -1 {
            minLen = length
        } else if minLen > length {
            minLen = length
        }
    }

    return minLen
}

func maxElement(list map[int]int) int {
    max := 0

    for _, v := range list {
        if max < v {
            max = v
        }
    }

    return max
}

func degreeList(nums []int) map[int]int {
    var count map[int]int

    maxDegree := 0
    count = make(map[int]int)

    for _, v := range nums {
        count[v]++
        if maxDegree < count[v] {
            maxDegree = count[v]
        }
    }

    return count
}
