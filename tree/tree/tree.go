package tree

import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Sample() *TreeNode {
    root := newTreeNode(3)
    root.Left = newTreeNode(9)
    root.Right = newTreeNode(20)
    root.Right.Left = newTreeNode(15)
    root.Right.Right = newTreeNode(7)

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
