/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    oldToNew := make(map[*Node]*Node)

    var dfs func(*Node) *Node
    dfs = func(node *Node) *Node {
        if node == nil {
            return nil
        }

        if _, found := oldToNew[node]; found {
            return oldToNew[node]
        }

        copy := &Node{Val: node.Val}
        oldToNew[node] = copy
        for _, nei := range node.Neighbors {
            copy.Neighbors = append(copy.Neighbors, dfs(nei))
        }
        return copy
    }

    return dfs(node)
}