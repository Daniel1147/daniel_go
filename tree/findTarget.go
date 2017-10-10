package tree

// import "fmt"
import . "github.com/Daniel1147/tree"

var nums []int

func findTarget(root *TreeNode, k int) bool {
    nums = make([]int, 0)
    buildSliceFromTree(root)
    remainMap := make(map[int]int)
    for index, value := range nums {
        remain := k - value
        _, exist := remainMap[remain]
        if (exist) {
            return true
        }
        remainMap[value] = index
    }

    return false
}

func buildSliceFromTree(root *TreeNode) {
    if (root == nil) {
        return
    }

    buildSliceFromTree(root.Left)
    nums = append(nums, root.Val)
    buildSliceFromTree(root.Right)
}
