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

    return symmetric(root.Left, root.Right)
}

func symmetric(left *TreeNode, right *TreeNode) bool {
    if (left == nil && right == nil) {
        return true
    }

    if (left == nil || right == nil) {
        return false
    }

    if (left.Val != right.Val) {
        return false
    }

    if (symmetric(left.Left, right.Right) != true) {
        return false
    }

    if (symmetric(left.Right, right.Left) != true) {
        return false
    }

    return true
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(isSymmetric(treeNode))

	return
}

