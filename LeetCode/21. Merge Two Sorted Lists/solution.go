/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode = nil
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp := list1.Next
			list1.Next = result
			result = list1
			list1 = temp
		} else {
			temp := list2.Next
			list2.Next = result
			result = list2
			list2 = temp
		}
	}
	for list1 != nil {
		temp := list1.Next
		list1.Next = result
		result = list1
		list1 = temp
	}
	for list2 != nil {
		temp := list2.Next
		list2.Next = result
		result = list2
		list2 = temp
	}
	var prev *ListNode = nil
	curr := result

	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}