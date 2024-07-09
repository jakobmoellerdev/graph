package main

// Node represents a node in a directed graph with its outgoing edges.
type Node struct {
	connectedTo []int
}

// Graph represents a directed graph of nodes with embedded edges.
// The graph is represented as an adjacency list.
// The following would  be an example graph:
// +---+
// | 0 |
// +---+
//
// +------------------------+
// |                        |
// |  +---+     +---+     +---+     +---+     +---+
// |  | 8 | --> | 4 | --> | 9 | --> | 6 | --> | 2 |
// |  +---+     +---+     +---+     +---+     +---+
// |    |                   ^
// |    +---------+         |
// |              v         |
// |  +---+     +---+     +---+
// |  | 1 | --> | 3 |     | 7 |
// |  +---+     +---+     +---+
// |    |
// |    |
// |    v
// |  +---+
// +> | 5 |
type Graph struct {
	nodes []Node
}

// NewGraph creates a new empty graph with an underlying slice with n capacity
func NewGraph(n int) *Graph {
	return nil
}

// Size returns the number of nodes in the graph
func (g *Graph) Size() int {
	return 0
}

// AddNode adds a new node to the graph
func (g *Graph) AddNode(node Node) {
}

// GetNode returns the node with id i
func (g *Graph) GetNode(i int) *Node {
	return nil
}

// AddEdge adds an edge between two nodes
func (g *Graph) AddEdge(start, end int) {

}

// GetOutgoingEdges returns the outgoing edges of a node
func (g *Graph) GetOutgoingEdges(node int) []int {
	return nil
}

// GetOutgoingEdges returns the incoming edges of a node
func (g *Graph) GetIncomingEdges(node int) []int {
	return nil
}

// GetNeighbours returns the neighbours of a node at a certain level.
// The 1st level neighbours are the direct neighbours of the node, the second level neighbours are the neighbours of the neighbours and so on.
func (g *Graph) GetNeighbours(node int, level int) []int {
	return nil
}

// BFS performs a breath-first traversal starting from a given node
// start is the node to start the traversal from.
// visited is a map that keeps track of the visited nodes.
// result is a slice that keeps track of the visited nodes in the order they were visited.
func (g *Graph) BFS(start, target int) []int {
	return nil
}
