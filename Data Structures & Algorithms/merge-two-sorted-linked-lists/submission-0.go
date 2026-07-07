/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    dummyNode := &ListNode{}
	tail := dummyNode // pointer to empty struct

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	
	// тепер обробка списку який закінчився, додаємо в кінець той що лишився
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummyNode.Next
}
