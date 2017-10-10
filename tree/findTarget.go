package tree

// import "fmt"
import . "github.com/Daniel1147/tree"

var remainMap map[int]int

func findTarget(root *TreeNode, k int) bool {
    remainMap = make(map[int]int)

    return findTargetHelper(root, k)
}

func findTargetHelper(node *TreeNode, target int) bool {
    if node == nil {
        return false
    }
    remain := target - node.Val
    _, exist := remainMap[node.Val]
    if (exist == true) {
        return true
    }

    remainMap[remain] = 1
    leftResult := findTargetHelper(node.Left, target)
    if leftResult == true {
        return true
    }

    rightResult := findTargetHelper(node.Right, target)
    if rightResult == true {
        return true
    }
    return false
}
