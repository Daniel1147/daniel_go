package main

import "fmt"
import . "./list"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if (l1 == nil) {
        return l2
    }

    if (l2 == nil) {
        return l1
    }

    if (l1.Val < l2.Val) {
        outNode := l1
        outNode.Next = mergeTwoLists(outNode.Next, l2)
        return outNode
    } else {
        outNode := l2
        outNode.Next = mergeTwoLists(outNode.Next, l1)
        return outNode
    }
}

func buildList(arr []ListNode) *ListNode {
    var lastNode *ListNode
    var firstNode *ListNode
    for nodeIndex := range arr {
        if (lastNode == nil) {
            lastNode = &arr[nodeIndex]
            firstNode = &arr[nodeIndex]
        } else {
            lastNode.Next = &arr[nodeIndex]
            lastNode = &arr[nodeIndex]
        }
    }

    return firstNode
}

func dumpList(node *ListNode) []string {
    for {
        if (node == nil) {
            break
        }
        fmt.Println(node.Val)
        node = node.Next
    }
    return []string{}
}

func main() {
    l1 := []ListNode{
        {
            5,
            nil,
        },
        {
            6,
            nil,
        },
        {
            7,
            nil,
        },
    }
    list1 := buildList(l1)

    l2 := []ListNode{
        // {
        //     1,
        //     nil,
        // },
        // {
        //     2,
        //     nil,
        // },
        // {
        //     3,
        //     nil,
        // },
    }
    list2 := buildList(l2)

    dumpList(mergeTwoLists(list1, list2))

	return
}

