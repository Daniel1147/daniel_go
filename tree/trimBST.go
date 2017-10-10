package main

// import "fmt"
import . "github.com/Daniel1147/tree"

func trimBST(root *TreeNode, L int, R int) *TreeNode {
    if root == nil {
        return nil
    }

    if (root.Val >= L && root.Val <= R) {
        root.Left = trimBST(root.Left, L, root.Val)
        root.Right = trimBST(root.Right, root.Val, R)
        return root
    }

    if (root.Val < L) {
        root = trimBST(root.Right, L, R)
        return root
    }

    if (root.Val > R) {
        root = trimBST(root.Left, L, R)
        return root
    }

    return root
}
