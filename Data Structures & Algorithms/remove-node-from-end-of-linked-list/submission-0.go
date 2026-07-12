/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    N := 0
    cur := head
    for cur != nil {
        N++
        cur = cur.Next
    }

    removeIndex := N - n
    if removeIndex == 0 {
        return head.Next
    }

    cur = head
    for i := 0; i < N-1; i++ {
        if (i + 1) == removeIndex {
            cur.Next = cur.Next.Next
            break
        }
        cur = cur.Next
    }
    return head
}