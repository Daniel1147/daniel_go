package list

import "fmt"
import . "github.com/Daniel1147/linkList"
import "testing"

func TestAverageOfLevels(t *testing.T) {
    cases := []struct {
        list1 *ListNode
        list2 *ListNode
        expected string
    } {
        sample1(),
        sample2(),
        sample3(),
    }

    for _, testCase := range cases {
        list1Dump := testCase.list1.Dump()
        list2Dump := testCase.list2.Dump()

        // get result
        resultSlice := addTwoNumbers(testCase.list1, testCase.list2).Dump()
        result := fmt.Sprintf("%v", resultSlice)

        // build expected
        expected := testCase.expected

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\nlist1 %v\nlist2 %v\nexpected %v\nresult %v",
                list1Dump,
                list2Dump,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() (result struct {
    list1 *ListNode
    list2 *ListNode
    expected string
}) {
    slice1 := []int{2, 4, 3}
    slice2 := []int{5, 6, 4}
    expectedSlice := []int{7, 0, 8}

    result.list1 = BuildListBySlice(slice1)
    result.list2 = BuildListBySlice(slice2)
    result.expected = fmt.Sprintf("%v", BuildListBySlice(expectedSlice).Dump())

    return
}

func sample2() (result struct {
    list1 *ListNode
    list2 *ListNode
    expected string
}) {
    slice1 := []int{0, 0, 1}
    slice2 := []int{1}
    expectedSlice := []int{1, 0, 1}

    result.list1 = BuildListBySlice(slice1)
    result.list2 = BuildListBySlice(slice2)
    result.expected = fmt.Sprintf("%v", BuildListBySlice(expectedSlice).Dump())

    return
}

func sample3() (result struct {
    list1 *ListNode
    list2 *ListNode
    expected string
}) {
    slice1 := []int{9, 9, 9}
    slice2 := []int{1}
    expectedSlice := []int{0, 0, 0, 1}

    result.list1 = BuildListBySlice(slice1)
    result.list2 = BuildListBySlice(slice2)
    result.expected = fmt.Sprintf("%v", BuildListBySlice(expectedSlice).Dump())

    return
}
