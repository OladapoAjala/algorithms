package dynamicprogramming

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/graph"
	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/constraints"
)

func Test_MatrixShortestPath(t *testing.T) {
	graph := graph.NewGraph[string, int]()
	_ = graph.Add(0, "", "A")
	_ = graph.Add(1, "A", "B")
	_ = graph.Add(1, "A", "D")
	_ = graph.Add(2, "B", "E")
	_ = graph.Add(1, "D", "E")

	testCases := []struct {
		Start        string
		Stop         string
		ExpectedPath []string
	}{
		{"A", "E", []string{"A", "D", "E"}},
	}

	for _, tc := range testCases {
		start := search(graph, tc.Start)
		stop := search(graph, tc.Stop)
		path := ShortestPath(start, stop)
		fmt.Println(path)
	}
}

func Test_ShortestPath(t *testing.T) {
	graph := graph.NewGraph[string, int]()
	graph.Add(0, "", "A")
	graph.Add(2, "A", "B")
	graph.Add(1, "A", "C")
	graph.Add(1, "B", "D")
	graph.Add(2, "B", "E")
	graph.Add(1, "C", "F")
	graph.Add(2, "C", "B")

	testCases := []struct {
		Start         string
		Stop          string
		ExpectedPath  []string
		ExpectedError error
	}{
		{"A", "B", []string{"A", "B"}, nil},
		{"A", "F", []string{"A", "C", "F"}, nil},
		// {"E", "A", nil, errors.New("no path from E -> A")},
		// {"X", "Y", nil, errors.New("data X not found in graph")},
		// {"A", "X", nil, errors.New("data X not found in graph")},
		{"A", "A", []string{"A"}, nil},
		{"A", "C", []string{"A", "C"}, nil},
		{"D", "D", []string{"D"}, nil},
		// {"F", "B", nil, errors.New("no path from F -> B")},
	}

	for _, tc := range testCases {
		start := search(graph, tc.Start)
		stop := search(graph, tc.Stop)
		path := ShortestPath(start, stop)
		require.Equal(t, tc.ExpectedPath, path)
	}
}

func search[V comparable, W constraints.Ordered](g *graph.Graph[V, W], data V) *vertex.Vertex[V, W] {
	for i, d := range g.Vertices {
		if d.State == data {
			return g.Vertices[i]
		}
	}
	return nil
}
