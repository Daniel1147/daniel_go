package main

import "fmt"
import . "github.com/Daniel1147/linkList"
import "testing"

func TestReverseList(t *testing.T) {
    cases := []struct {
        root *ListNode
        expectedSlice []int
    } {
        sample1(),
        sample2(),
        sample3(),
    }

    for _, testCase := range cases {
        caseDump := testCase.root.Dump()
        // get result
        resultSlice := reverseList(testCase.root).Dump()
        result := fmt.Sprintf("%v", resultSlice)

        // build expected
        expected := fmt.Sprintf("%v", testCase.expectedSlice)

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\nexpected %v\nresult %v",
                caseDump,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() (result struct {
    root *ListNode
    expectedSlice []int
}) {
    rawList := []int{}
    result.root = BuildListBySlice(rawList)
    result.expectedSlice = []int{}
    return
}

func sample2() (result struct {
    root *ListNode
    expectedSlice []int
}) {
    rawList := []int{1}
    result.root = BuildListBySlice(rawList)
    result.expectedSlice = []int{1}
    return
}

func sample3() (result struct {
    root *ListNode
    expectedSlice []int
}) {
    rawList := []int{1, 2}
    result.root = BuildListBySlice(rawList)
    result.expectedSlice = []int{2, 1}
    return
}
