package tree

// import "fmt"
import . "github.com/Daniel1147/tree"

func hasPathSum(root *TreeNode, sum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        if (root.Val == sum) {
            return true
        }
        return false
    }

    if hasPathSum(root.Left, sum - root.Val) == true {
        return true
    }

    if hasPathSum(root.Right, sum - root.Val) == true {
        return true
    }

    return false
}
