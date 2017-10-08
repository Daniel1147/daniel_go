package main

import "fmt"
import . "./myApp"

func levelOrderBottom(root *TreeNode) [][]int {
    var result, leftResult, rightResult [][]int
    if root == nil {
        return result
    }
    if (root.Left != nil) {
        leftResult = levelOrderBottom(root.Left)
    }

    if (root.Right != nil) {
        rightResult = levelOrderBottom(root.Right)
    }

    result = myMerge(leftResult, rightResult)

    // add a layer
    result = append(result, []int{root.Val})
    return result
}

func myMerge(s1, s2 [][]int) [][]int {
    // fmt.Println("merge", s1, s2)
    var result [][]int
    s1Len := len(s1)
    s2Len := len(s2)
    var longLen int

    if (s1Len > s2Len) {
        longLen = s1Len
    } else {
        longLen = s2Len
    }

    for i := 0; i < longLen; i++ {
        var currentLevel []int
        i1 := s1Len - longLen + i
        if (i1 >= 0) {
            currentLevel = append(currentLevel, s1[i1]...)
        }
        i2 := s2Len - longLen + i
        if (i2 >= 0) {
            currentLevel = append(currentLevel, s2[i2]...)
        }
        result = append(result, currentLevel)
    }

    // fmt.Println("result", result)
    return result
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(levelOrderBottom(treeNode))

	return
}

