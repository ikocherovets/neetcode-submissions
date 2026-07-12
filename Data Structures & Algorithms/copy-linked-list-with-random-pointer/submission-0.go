/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

var m = make(map[*Node]*Node)

func copyRandomList(head *Node) *Node {
    if head == nil {
        return nil
    }
    if val, exists := m[head]; exists {
        return val
    }

    copy := &Node{Val: head.Val}
    m[head] = copy
    copy.Next = copyRandomList(head.Next)
    copy.Random = m[head.Random]
    return copy
}