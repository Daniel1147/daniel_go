package main

import . "github.com/Daniel1147/tree"

func sortedArrayToBST(nums []int) *TreeNode {
    var root *TreeNode
    if len(nums) == 0 {
        return root
    }
    if len(nums) == 1 {
        root = buildNode(nums[0])
        return root
    }
    middleIndex := len(nums) / 2
    root = buildNode(nums[middleIndex])
    root.Left = sortedArrayToBST(nums[:middleIndex])
    root.Right = sortedArrayToBST(nums[middleIndex + 1:])

    return root
}

func buildNode(num int) *TreeNode {
    newNode := &TreeNode{num, nil, nil}

    return newNode
}

func insertNode(root *TreeNode, node *TreeNode) {
    if root == nil {
        root = node
    }

    if (node.Val <= root.Val) {
        if root.Left == nil {
            root.Left = node
        } else {
            insertNode(root.Left, node)
        }
    } else {
        if root.Right == nil {
            root.Right = node
        } else {
            insertNode(root.Right, node)
        }
    }
    return
}

func main() {
    input := []int{3, 5, 8}
    result := sortedArrayToBST(input)
    result.Show()

	return
}

