package main

import "fmt"
import . "./myApp"

func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if root.Left == nil && root.Right == nil {
        return true
    } else if root.Left == nil || root.Right == nil {
        return false
    }
    leftDump := dump(root.Left, 'L')
    rightDump := dump(root.Right, 'R')

    return sliceEqual(leftDump, rightDump)
}

func sliceEqual(s1, s2 []int) bool {
    if len(s1) != len(s2) {
        return false
    }

    for i := 0; i < len(s1); i++ {
        if (s1[i] != s2[i]) {
            return false
        }
    }
    return true
}

func dump(node *TreeNode, direction byte) []int {
    var left, right, middle []int
    if (node.Left != nil) {
        left = dump(node.Left, direction)
    } else {
        left = []int{-1}
    }

    middle = []int{node.Val}

    if (node.Right != nil) {
        right = dump(node.Right, direction)
    } else {
        right = []int{-1}
    }

    if direction == 'L' {
        return append(middle, append(left, right...)...)
    } else {
        return append(middle, append(right, left...)...)
    }
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(isSymmetric(treeNode))

	return
}

