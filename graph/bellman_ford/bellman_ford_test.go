package bellmanford

import (
	"fmt"
	"testing"

	"github.com/OladapoAjala/datastructures/graph"
	"github.com/Pallinder/go-randomdata"
	"github.com/stretchr/testify/assert"
)

func GenerateGridGraph(width, height int) *graph.Graph[string, int] {
	g := graph.NewGraph[string, int]()

	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			vertexID := fmt.Sprintf("%d%d", i, j)
			weight := randomdata.Number(-5, 5)
			_ = g.Add(weight, vertexID, fmt.Sprintf("%d%d", i+1, j))
			weight = randomdata.Number(-5, 5)
			_ = g.Add(weight, vertexID, fmt.Sprintf("%d%d", i, j+1))
		}
	}
	return g
}

func Test_bellmanFord(t *testing.T) {
	is := assert.New(t)
	testGraph := GenerateGridGraph(10, 10)

	testCases := []struct {
		Start         string
		Stop          string
		ExpectedPath  []string
		ExpectedError error
	}{
		{"00", "59", []string{"A", "B"}, nil},
	}

	for _, tc := range testCases {
		start, err := testGraph.Search(tc.Start)
		is.Nil(err)
		end, err := testGraph.Search(tc.Stop)
		is.Nil(err)

		path, err := bellmanFord(testGraph, start, end)
		is.Nil(err)
		fmt.Println(path)
	}
}
