package uncategorized

import "testing"
import "fmt"

type TestCase struct {
    input string
    expected bool
}

func TestIsPalindrome(t *testing.T) {
    testList := []TestCase{
        case1(),
        case2(),
        case3(),
    }

    for _, testCase := range testList {
        // get result
        result := isPalindrome(testCase.input)

        // build expected
        expected := testCase.expected

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\nexpected %v\nresult %v",
                testCase.input,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func case1() (t TestCase) {
    t.input = "A man, a plan, a canal: Panama"
    t.expected = true
    return
}

func case2() (t TestCase) {
    t.input = "race a car"
    t.expected = false
    return
}

func case3() (t TestCase) {
    t.input = "010"
    t.expected = true
    return
}

func case4() (t TestCase) {
    t.input = "01"
    t.expected = false
    return
}
