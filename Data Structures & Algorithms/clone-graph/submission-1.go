/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	oldNewNodeMap := make(map[*Node]*Node)
    var dfs func(node *Node) *Node

	dfs = func(node *Node) *Node {
		// base case 
		// if node is nil return nil
		if node == nil {
			return nil
		}

		// далі перевіряємо чи вже маємо таку ноду в мапі
		if _, ok := oldNewNodeMap[node]; ok {
			return oldNewNodeMap[node]
		}

		clone := &Node{Val: node.Val}
		oldNewNodeMap[node] = clone

		// а тепер faith of loop
		// клонуємо всіх сусідів
		for _, neighbour := range node.Neighbors {
			clone.Neighbors = append(clone.Neighbors, dfs(neighbour))
		}

		return clone
	}

	return dfs(node)
}
