package main

import . "github.com/Daniel1147/tree"

func sortedListToBST(head *ListNode) *TreeNode {
    sortedArray := listToArray(head)
    result := sortedArrayToBST(sortedArray)

    return result
}

func listToArray(root *ListNode) []int {
    var result []int
    lastNode := root
    for lastNode != nil {
        result = append(result, lastNode.Val)
        lastNode = lastNode.Next
    }
    return result
}

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
