package tree

import "fmt"
import . "github.com/Daniel1147/tree"

type levelData struct {
    num int
    sum int
}

var levelDataList []levelData

func averageOfLevels(root *TreeNode) []float64 {
    levelDataList = make([]levelData, 0)
    traverse(root, 0)
    fmt.Println(levelDataList)
    result := buildResult()
    return result
}

func traverse(node *TreeNode, level int) {
    if (node == nil) {
        return
    }

    currentLevel := level + 1

    // TODO
    // add data
    if len(levelDataList) < currentLevel {
        levelDataList = append(levelDataList, levelData{0, 0})
    }
    levelDataList[level].num += 1
    levelDataList[level].sum += node.Val

    traverse(node.Left, currentLevel)
    traverse(node.Right, currentLevel)
    return
}

func buildResult() []float64 {
    var result []float64
    for _, v := range levelDataList {
        result = append(result, float64(v.sum) / float64(v.num))
    }
    return result
}
