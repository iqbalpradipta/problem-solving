package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val	int
	Next *ListNode
}


func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := &ListNode{}
	num := res

	for list1 != nil || list2 != nil {
		if list1 != nil {
			num.Val += list1.Val
			list1 = list1.Next
		}

		if list2 != nil {
			num.Val += list2.Val
			list2 = list2.Next
		}

		if num.Val > 9 {
			num.Val -= 10
			num.Next = &ListNode{Val: 1}
		} else if list1 != nil || list2 != nil {
			num.Next = &ListNode{}
		}

		num = num.Next
	}
	return res
}