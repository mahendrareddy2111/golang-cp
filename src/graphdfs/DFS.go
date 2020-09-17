package graphdfs

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	// Write your code here.
	if n == nil {
		return array
	}

	if n != nil && n.Children == nil {
		array = append(array,n.Name)
		return array
	}
	array = append(array,n.Name)
	for _,c := range n.Children{
		array = c.DepthFirstSearch(array)
	}
	return array
}