package greedy

import (
	"fmt"
	"testing"
)

type TestCase struct {
	children []int
	cookie   []int
	expected int
}

func TestFindContentChildren(t *testing.T) {
	testCaseList := []TestCase{
		case1(),
		case2(),
		case3(),
		case4(),
	}

	for _, testCase := range testCaseList {
		// input
		children := fmt.Sprintf("%v", testCase.children)
		cookie := fmt.Sprintf("%v", testCase.cookie)

		// result
		result := findContentChildren(testCase.children, testCase.cookie)

		// expected
		expected := testCase.expected

		if expected != result {
			// error msg
			errorMsg := fmt.Sprintf(
				"\nchildren: %v\ncookie: %v\nresult: %v\nexpected: %v\n",
				children,
				cookie,
				result,
				expected,
			)

			t.Errorf(errorMsg)
		}
	}
}

func case1() (t TestCase) {
	t.children = []int{}
	t.cookie = []int{}
	t.expected = 0
	return
}

func case2() (t TestCase) {
	t.children = []int{1, 2, 3}
	t.cookie = []int{1, 1}
	t.expected = 1
	return
}

func case3() (t TestCase) {
	t.children = []int{1, 2}
	t.cookie = []int{1, 2, 3}
	t.expected = 2
	return
}

func case4() (t TestCase) {
	t.children = []int{10, 9, 8, 7}
	t.cookie = []int{5, 6, 7, 8}
	t.expected = 2
	return
}
