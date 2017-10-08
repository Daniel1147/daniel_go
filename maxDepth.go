package main

import "fmt"
import . "./myApp"

func maxDepth(root *TreeNode) int {
    if (root == nil) {
        return 0;
    }

    leftDepth := maxDepth(root.Left) + 1
    rightDepth := maxDepth(root.Right) + 1

    if (leftDepth > rightDepth) {
        return leftDepth
    } else {
        return rightDepth
    }
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(maxDepth(treeNode))

	return
}

