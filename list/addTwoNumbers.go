package list

// import "fmt"
import . "github.com/Daniel1147/linkList"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    // reverse1 := reverseList(l1)
    // reverse2 := reverseList(l2)
    // result := addTwoReversedList(reverse1, reverse2)
    result := addTwoList(l1, l2)
    return result
}

func addTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
    var lastNode, newNode, root *ListNode
    var currentVal, nextVal int
    for l1 != nil || l2 != nil {
        if l1 == nil {
            currentVal = l2.Val + nextVal
            l2 = l2.Next
        } else if (l2 == nil) {
            currentVal = l1.Val + nextVal
            l1 = l1.Next
        } else {
            currentVal = l1.Val + l2.Val + nextVal
            l1 = l1.Next
            l2 = l2.Next
        }
        nextVal = currentVal / 10
        currentVal = currentVal % 10
        // newNode = newListNode(currentVal)
        newNode = &ListNode{currentVal, nil}

        if root == nil {
            root = newNode
            lastNode = root
        } else {
            lastNode.Next = newNode
            lastNode = newNode
        }
    }
    if nextVal != 0 {
        // newNode = newListNode(nextVal)
        newNode = &ListNode{nextVal, nil}
        lastNode.Next = newNode
        lastNode = newNode
    }

    return root
}
