package dijkstra

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/graph"
	"github.com/stretchr/testify/assert"
)

func GenerateGridGraph(width, height int) *graph.Graph[string, int] {
	g := graph.NewGraph[string, int]()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			vertexID := fmt.Sprintf("%d%d", i, j)
			_ = g.Add(1, vertexID, fmt.Sprintf("%d%d", i+1, j))
			_ = g.Add(1, vertexID, fmt.Sprintf("%d%d", i, j+1))
		}
	}
	return g
}

func BenchmarkDijkstraOnGrid(b *testing.B) {
	is := assert.New(b)
	testGraph := GenerateGridGraph(100, 100)

	start, err := testGraph.Search("A")
	is.Nil(err)
	end, err := testGraph.Search("B")
	is.Nil(err)

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := dijkstra(testGraph, start, end)
		if err != nil {
			b.Fatalf("Error running Dijkstra: %v", err)
		}
	}
}

func Test_dikstra(t *testing.T) {
	is := assert.New(t)

	testGraph := GenerateGridGraph(10, 10)

	testCases := []struct {
		Start         string
		Stop          string
		ExpectedPath  []string
		ExpectedError error
	}{
		{"00", "97", []string{"A", "B"}, nil}, //"1318"
	}

	for _, tc := range testCases {
		start, err := testGraph.Search(tc.Start)
		is.Nil(err)
		end, err := testGraph.Search(tc.Stop)
		is.Nil(err)

		path, err := dijkstra(testGraph, start, end)
		is.Nil(err)
		fmt.Println(path)
	}
}
