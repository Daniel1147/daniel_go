package myApp

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Sample() *TreeNode {
    root := newTreeNode(1)
    root.Left = newTreeNode(2)
    // root.Left.Left = newTreeNode(3)
    root.Right = newTreeNode(4)

    return root
}

func newTreeNode(n int) * TreeNode {
    newNode := TreeNode{
        n,
        nil,
        nil,
    }

    return &newNode
}

func (node * TreeNode) Show () {
    if (node == nil) {
        return
    }

    fmt.Println(node.Val)
    node.Left.Show()
    node.Right.Show()
    return
}
