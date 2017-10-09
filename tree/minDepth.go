package main

import "fmt"
import . "github.com/Daniel1147/tree"

func minDepth(root *TreeNode) int {
    depth := 0

    if (root == nil) {
        return depth
    }
    currentList := []*TreeNode{root}
    currentListLen := 1
    depth++
    for {
        newList := make([]*TreeNode, currentListLen * 2)
        currentListLen = 0
        for _, node := range currentList {
            if node == nil {
                break
            }
            if node.Left == nil && node.Right == nil {
                return depth
            }

            if (node.Left != nil) {
                // newList = append(newList, node.Left)
                newList[currentListLen] = node.Left
                currentListLen++
            }

            if (node.Right != nil) {
                newList[currentListLen] = node.Right
                currentListLen++
            }
        }
        currentList = newList
        depth++
    }

    return depth
}
