package main

import "fmt"
import . "github.com/Daniel1147/tree"

func minDepth(root *TreeNode) int {
    depth := 0

    if (root == nil) {
        return depth
    }
    currentList := []*TreeNode{root}
    depth++
    for {
        newList := []*TreeNode{}
        for _, node := range currentList {
            if node.Left == nil && node.Right == nil {
                return depth
            }

            if (node.Left != nil) {
                newList = append(newList, node.Left)
            }

            if (node.Right != nil) {
                newList = append(newList, node.Right)
            }
        }
        currentList = newList
        depth++
    }

    return depth
}
