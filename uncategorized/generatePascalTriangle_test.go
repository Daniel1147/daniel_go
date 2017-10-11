package uncategorized

import "fmt"
import "testing"

func TestGeneratePascalTriangle(t *testing.T) {
    cases := []struct {
        num int
        expected [][]int
    } {
        sample1(),
        sample2(),
    }

    for _, testCase := range cases {
        // get result
        result := fmt.Sprintf("%v", generatePascalTriangle(testCase.num))

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
    expected [][]int
}) {
    result.num = 1
    result.expected = [][]int{
        []int{1},
    }
    return
}

func sample2() (result struct {
    num int
    expected [][]int
}) {
    result.num = 3
    result.expected = [][]int{
        []int{1},
        []int{1, 1},
        []int{1, 2, 1},
    }
    return
}
