package tree

// import "fmt"
import . "github.com/Daniel1147/tree"

var nillNum = -1147

func findTarget(root *TreeNode, k int) bool {
    if root == nil {
        return false
    }
    result := searchBST(root, k - root.Val, root.Val)
    if result == true {
        return true
    }

    leftResult := findTargetHelper(root.Left, k, root)
    if leftResult == true {
        return true
    }

    rightResult := findTargetHelper(root.Right, k, root)
    if rightResult == true {
        return true
    }

    return false
}

func findTargetHelper(node *TreeNode, k int, root *TreeNode) bool {
    if node == nil {
        return false
    }
    result := searchBST(root, k - node.Val, node.Val)
    if result == true {
        return true
    }

    leftResult := findTargetHelper(node.Left, k, root)
    if leftResult == true {
        return true
    }

    rightResult := findTargetHelper(node.Right, k, root)
    if rightResult == true {
        return true
    }

    return false
}

func searchBST(root *TreeNode, targetNum int, usedNum int) bool {
    if (root == nil) {
        return false
    }

    if root.Val == usedNum {
        usedNum = nillNum
    } else {
        if root.Val == targetNum {
            return true
        }
    }

    leftResult := searchBST(root.Left, targetNum, usedNum)
    rightResult := searchBST(root.Right, targetNum, usedNum)

    if leftResult == true || rightResult == true {
        return true
    }

    return false
}
