/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
    if head == nil {
        return
    }

    nodes := []*ListNode{}
    cur := head
    for cur != nil {
        nodes = append(nodes, cur)
        cur = cur.Next
    }

    i, j := 0, len(nodes)-1
    for i < j {
        nodes[i].Next = nodes[j]
        i++
        if i >= j {
            break
        }
        nodes[j].Next = nodes[i]
        j--
    }

    nodes[i].Next = nil
}