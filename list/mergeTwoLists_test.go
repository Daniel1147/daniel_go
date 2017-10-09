package main

import . "github.com/Daniel1147/linkList"
import "testing"

func TestMergeTwoLists(t *testing.T) {
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

    l3 := []ListNode{
        {
            1,
            nil,
        },
        {
            2,
            nil,
        },
        {
            3,
            nil,
        },
    }
    list3 := buildList(l3)

    cases := []struct {
        list1 *ListNode
        list2 *ListNode
        expected []int
    } {
        {
            list1,
            list2,
            []int{5, 6, 7},
        },
        {
            list1,
            list3,
            []int{1, 2, 3, 5, 6, 7},
        },
    }

    for _, v := range cases {
        result := mergeTwoLists(v.list1, v.list2)
        resultSlice := result.Dump()
        for resultIndex, entry := range v.expected {
            if entry != resultSlice[resultIndex] {
                // print case
                t.Errorf(
                    "\nl1 %v\nl2 %v\nexpected %v\nget %v\n",
                    v.list1.Dump(),
                    v.list2.Dump(),
                    v.expected,
                    resultSlice,
                )
            }
        }
    }

	return
}

