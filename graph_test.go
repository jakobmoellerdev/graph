package main

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestNewGraphAndSize(t *testing.T) {
	var (
		expectedSize = 0
		capazity     = 10
		graph        = NewGraph(capazity)
	)

	if graph == nil {
		t.Fatal("graph should not be nil")
	}

	if cap(graph.nodes) != capazity {
		t.Errorf("graph does not have expected capazity (got %d want %d)", cap(graph.nodes), capazity)
	}

	if size := graph.Size(); size != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", size, expectedSize)
	}
}

func TestAddAndSize(t *testing.T) {
	var (
		expectedSize = 0
		graph        = NewGraph(10)
	)

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", capacity, expectedSize)
	}

	graph.AddNode(Node{})
	expectedSize++

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", capacity, expectedSize)
	}
}

func TestGet(t *testing.T) {
	var (
		expectedSize = 0
		graph        = NewGraph(10)
	)

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", capacity, expectedSize)
	}

	graph.AddNode(Node{})
	expectedSize++

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("tree does not have expected size (got %d want %d)", capacity, expectedSize)
	}
	if node := graph.GetNode(0); node == nil {
		t.Errorf("node sould not be nil")
	}
}

func TestAddEgde(t *testing.T) {
	var (
		expectedSize = 0
		graph        = NewGraph(10)
	)

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", capacity, expectedSize)
	}

	graph.AddNode(Node{})
	expectedSize++

	if capacity := graph.Size(); capacity != expectedSize {
		t.Errorf("graph does not have expected size (got %d want %d)", capacity, expectedSize)
	}

	graph.AddEdge(0, 1)

	if node := graph.GetNode(0); node == nil {
		t.Errorf("node sould not be nil")
	} else {
		if node.connectedTo[0] != 1 {
			t.Errorf("node should be connected to 1")
		}
	}
}

func TestOutgoingEdges(t *testing.T) {
	var (
		outgoingEdges = map[int][]int{
			0: {},
			1: {3, 5},
			2: {},
			3: {},
			4: {9},
			5: {},
			6: {2},
			7: {9},
			8: {3, 4},
			9: {5, 6},
		}
		graph = buildTestGraph()
	)

	for node, expectedOutgoingEdges := range outgoingEdges {
		if outgoingEdges := graph.GetOutgoingEdges(node); !reflect.DeepEqual(outgoingEdges, expectedOutgoingEdges) {
			t.Errorf("outgoing edges for %v does not have expected (got %v want %v)", node, outgoingEdges, expectedOutgoingEdges)
		}
	}
}

func TestIncomingEdges(t *testing.T) {
	var (
		incomingEdges = map[int][]int{
			0: {},
			1: {},
			2: {6},
			3: {1, 8},
			4: {8},
			5: {1, 9},
			6: {9},
			7: {},
			8: {},
			9: {4, 7},
		}
		graph = buildTestGraph()
	)

	for node, expectedIncomingEdges := range incomingEdges {
		if incomingEdges := graph.GetIncomingEdges(node); !reflect.DeepEqual(incomingEdges, expectedIncomingEdges) {
			t.Errorf("incoming edges for %v does not have expected (got %v want %v)", node, incomingEdges, expectedIncomingEdges)
		}
	}
}

func TestNeighboors(t *testing.T) {
	var (
		neighboorsForGrade = map[int][]int{
			1: {3, 4},
			2: {9},
			3: {5, 6},
		}
		graph = buildTestGraph()
	)

	for nLevel, expectedNeighboors := range neighboorsForGrade {
		if neighboors := graph.GetNeighbours(8, nLevel); !reflect.DeepEqual(neighboors, expectedNeighboors) {
			t.Errorf("neighboors %v level does not have expected (got %v want %v)", nLevel, neighboors, expectedNeighboors)
		}
	}

}

func TestBFS(t *testing.T) {
	graph := buildTestGraph()

	for _, tc := range []struct {
		start, end int
		path       []int
	}{
		{8, 2, []int{8, 4, 9, 6, 2}},
		{9, 2, []int{9, 6, 2}},
		{1, 5, []int{1, 5}},
		{3, 9, []int{}},
	} {
		t.Run("", func(t *testing.T) {
			result := graph.BFS(tc.start, tc.end)

			if !reflect.DeepEqual(result, tc.path) {
				t.Errorf("path does not have expected (got %v want %v)", result, tc.path)
			}
		})
	}

}

// +---+
// | 0 |
// +---+

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
func buildTestGraph() *Graph {
	var graph = NewGraph(10)

	for range 10 {
		graph.AddNode(Node{
			connectedTo: make([]int, 0),
		})
	}
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 5)
	graph.AddEdge(4, 9)
	graph.AddEdge(6, 2)
	graph.AddEdge(7, 9)
	graph.AddEdge(8, 3)
	graph.AddEdge(8, 4)
	graph.AddEdge(9, 5)
	graph.AddEdge(9, 6)

	return graph
}
