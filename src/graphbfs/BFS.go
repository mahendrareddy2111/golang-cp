package graphbfs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	// Write your code here.
	if n == nil {
		return array
	}
	tempArr := []*Node{n}

	for len(tempArr) != 0 {
		node := tempArr[0]
		array = append(array, node.Name)
		for _, val := range node.Children {
			tempArr = append(tempArr, val)
		}
		tempArr = tempArr[1:]
	}

	return array
}
