package main

// import "fmt"
import . "github.com/Daniel1147/linkList"

func reverseList(head *ListNode) *ListNode {
    var current, next *ListNode
    if (head == nil) {
        return current
    }

    current = head
    next = current.Next
    current.Next = nil
    for next != nil {
        newNext := next.Next
        next.Next = current
        current = next
        next = newNext
    }

    return current
}
