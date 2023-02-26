package codetest

type ListNode struct {
	Val 	int
	Next	*ListNode
}

func reverseList(head *ListNode) *ListNode {
    var prev, cur *ListNode = nil, head

	for cur != nil {
		prev, cur, cur.Next = cur, cur.Next, prev
	}
	return prev
}