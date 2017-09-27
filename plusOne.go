package main

import "fmt"

func plusOne(digits []int) []int {
    ans := make([]int, len(digits) + 1)

    over := 1
    for i := len(digits) - 1; i >= 0; i-- {
        num := digits[i] + over
        if (num == 10) {
            over = 1
            ans[i+1] = 0
        } else {
            over = 0
            ans[i+1] = num
        }
    }
    ans[0] = over

    if (over == 1) {
        return ans
    } else {
        return ans[1:]
    }
}

func plusOneP(digits []int) []int {
    over := 1
    for i := len(digits) - 1; i >= 0; i-- {
        if (over == 0) {
            return digits
        }
        num := digits[i] + over
        if (num == 10) {
            over = 1
            digits[i] = 0
        } else {
            over = 0
            digits[i] = num
        }
    }

    if (over == 1) {
        return append([]int{1}, digits...)
    } else {
        return digits
    }
}

func plusOneP2(digits []int) []int {
    for i := len(digits) - 1; i >= 0; i-- {
        if (digits[i] != 9) {
            digits[i] += 1
            return digits
        }
        digits[i] = 0
    }

    ans := make([]int, len(digits) + 1);
    ans[0] = 1
    copy(ans[1:], digits)

    return ans
}

func main() {
    input1 := []int{1, 2, 3, 4}
    input2 := []int{9}

    fmt.Println(plusOneP2(input1))
    fmt.Println(plusOneP2(input2))

	return
}

