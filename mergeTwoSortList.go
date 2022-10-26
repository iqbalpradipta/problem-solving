package problemsolving

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res :=  &ListNode{}
	resTemp := res

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			resTemp.Next = list1
			list1 = list1.Next
		} else {
			resTemp.Next = list2
			list2 = list2.Next	
		}
		resTemp = resTemp.Next
		if list1 == nil {
			resTemp.Next = list2
		}else {
			resTemp.Next = list1
		}
	}
	return res.Next
}