package uncategorized

import "fmt"
import "testing"

func TestMaxProfit2(t *testing.T) {
    cases := []struct {
        num []int
        expected int
    } {
        sample1(),
        sample2(),
        sample3(),
        sample4(),
    }

    for _, testCase := range cases {
        // get result
        result := maxProfit2(testCase.num)

        // build expected
        expected := testCase.expected

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\nexpected %v\nresult %v",
                testCase.num,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() (result struct {
    num []int
    expected int
}) {
    result.num = []int{1}
    result.expected = 0

    return
}

func sample2() (result struct {
    num []int
    expected int
}) {
    result.num = []int{}
    result.expected = 0

    return
}

func sample3() (result struct {
    num []int
    expected int
}) {
    result.num = []int{1, 2, 3, 1, 2, 3}
    result.expected = 4

    return
}

func sample4() (result struct {
    num []int
    expected int
}) {
    result.num = []int{3, 2, 1, 3, 2, 1}
    result.expected = 2

    return
}
