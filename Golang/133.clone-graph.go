package golang

type Node struct {
	Val       int
	Neighbors []*Node
}

/*
 * @lc app=leetcode id=133 lang=golang
 *
 * [133] Clone Graph
 */

// @lc code=start
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

	stack := []*Node{node}
	var copied = make(map[int]*Node)
	var visited = map[int]bool{}

	// copy nodes with val, bfs
	for len(stack) > 0 {
		temp := stack[0]
		stack = stack[1:]

		if visited[temp.Val] {
			continue
		}

		visited[temp.Val] = true
		copied[temp.Val] = &Node{
			Val: temp.Val,
		}

		stack = append(stack, temp.Neighbors...)
	}

	// copy the neighbors
	stack = []*Node{node}
	visited = map[int]bool{}
	for len(stack) > 0 {
		temp := stack[0]
		stack = stack[1:]

		if visited[temp.Val] {
			continue
		}

		visited[temp.Val] = true
		for _, neighbor := range temp.Neighbors {
			copied[temp.Val].Neighbors = append(copied[temp.Val].Neighbors, copied[neighbor.Val])
		}

		stack = append(stack, temp.Neighbors...)
	}

	return copied[node.Val]
}

// @lc code=end
