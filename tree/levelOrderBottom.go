package main

import "fmt"
import . "./tree"

var ansList [][]int

func levelOrderBottom(root *TreeNode) [][]int {
    ansList = make([][]int, 0)
    addNode(root, 0)

    // return reversed ansList
    ansListLen := len(ansList)
    result := make([][]int, len(ansList))
    for i := range ansList {
        result[i] = ansList[ansListLen - i - 1]
    }

    return result
}

func addNode(node *TreeNode, level int) {
    if node == nil {
        return
    }

    if len(ansList) <= level {
        ansList = append(ansList, []int{})
    }

    addNode(node.Left, level + 1)
    addNode(node.Right, level + 1)

    ansList[level] = append(ansList[level], node.Val)

    return
}

func main() {
    treeNode := Sample()
    treeNode.Show()
    fmt.Println(levelOrderBottom(treeNode))

	return
}

