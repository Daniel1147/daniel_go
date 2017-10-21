package array

import "testing"
import "fmt"

type TestCase struct {
    nums []int
    expected int
}

func TestFindShortestSubArray(t *testing.T) {
    testList := []TestCase{
        case1(),
        case2(),
    }

    for _, testCase := range testList {
        expected := testCase.expected
        result := findShortestSubArray(testCase.nums)

        if (result != expected) {
            errorMsg := fmt.Sprintf("\nresult: %v\nexpected: %v\n", result, expected)
            t.Errorf(errorMsg)
        }
    }
}

func case1() TestCase {
    testCase := TestCase{
        []int{1, 2, 2, 3, 1},
        2,
    }

    return testCase
}

func case2() TestCase {
    testCase := TestCase{
        []int{1, 2, 2, 3, 1, 4, 2},
        6,
    }

    return testCase
}
