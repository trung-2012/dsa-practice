/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode = nil
	curr := slow.Next
	slow.Next = nil
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	first := head
	second := prev

	for second != nil {
		temp1 := first.Next
		temp2 := second.Next
		first.Next = second
		second.Next = temp1
		first = temp1
		second = temp2
	}
}