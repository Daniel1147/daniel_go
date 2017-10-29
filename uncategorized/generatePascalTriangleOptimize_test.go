package uncategorized

import "fmt"
import "testing"

func TestGeneratePascalTriangleOptimize(t *testing.T) {
    cases := []struct {
        num int
        expected []int
    } {
        sample1(),
        sample2(),
    }

    for _, testCase := range cases {
        // get result
        result := fmt.Sprintf("%v", getRow(testCase.num))

        // build expected
        expected := fmt.Sprintf("%v", testCase.expected)

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
    num int
    expected []int
}) {
    result.num = 0
    result.expected = []int{1}
    return
}

func sample2() (result struct {
    num int
    expected []int
}) {
    result.num = 2
    result.expected = []int{1, 2, 1}
    return
}
