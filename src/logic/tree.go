package logic

// Node represents a game tree node.
type Node struct {
	// Value is the value of this node.
	Value [3][3]int

	// Children is a slice of children of this node.
	Children []*Node
}

// previous is the previous value inserted into children.
var previous int = X

// Tree generates a new game tree for tic-tac-toe.
func Tree() {
	var tree map[string]*Node = make(map[string]*Node)
	tree["root"] = &Node{
		Value:    [3][3]int{},
		Children: []*Node{},
	}
	generateRootChildren(tree, tree["root"].Value)
}

// generateRootChildren generates all children of the root node and recursively calls generateChildren for each of the root's children.
func generateRootChildren(tree map[string]*Node, mtx [3][3]int) {
	var child int
	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			var matrix [3][3]int = mtx
			if matrix[i][k] == EMPTY {
				matrix[i][k] = X
			}
			tree["root"].Children = append(tree["root"].Children, &Node{
				Value:    matrix,
				Children: []*Node{},
			})
			generateChildren(tree["root"].Children[child])
			child++
		}
	}
}

// generateChildren generates all children for a child node.
func generateChildren(node *Node) {
	var (
		child  int
		insert int = X
	)

	if previous == X {
		insert = O
	}

	for i := 0; i < 3; i++ {
		for k := 0; k < 3; k++ {
			var matrix [3][3]int = node.Value
			if matrix[i][k] == EMPTY {
				matrix[i][k] = insert
				previous = insert
			}
			if matrix != node.Value {
				node.Children = append(node.Children, &Node{
					Value:    matrix,
					Children: []*Node{},
				})
				generateChildren(node.Children[child])
				child++
			}
		}
	}
}
