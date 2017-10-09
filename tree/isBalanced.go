package main

// import "fmt"
import . "github.com/Daniel1147/tree"

func isBalanced(root *TreeNode) bool {
    balanced, _ := myBalanced(root)

    return balanced
}

func myBalanced(node *TreeNode) (balanced bool, maxLen int) {
    if node == nil {
        balanced = true
        maxLen = 0
        return
    }

    leftBalanced, leftLen := myBalanced(node.Left)
    if leftBalanced == false {
        balanced = false
        maxLen = 0
        return
    }

    rightBalanced, rightLen := myBalanced(node.Right)
    if rightBalanced == false {
        balanced = false
        maxLen = 0
        return
    }

    balanced = balancedLen(leftLen, rightLen)
    if (leftLen > rightLen) {
        maxLen = leftLen + 1
        return
    }
    maxLen = rightLen + 1
    return
}

func balancedLen(len1, len2 int) bool {
    diff := len2 - len1
    if (diff * diff > 1) {
        return false
    }

    return true
}
