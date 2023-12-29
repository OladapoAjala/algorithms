package bellmanford

import (
	"github.com/OladapoAjala/datastructures/graph"
	"github.com/OladapoAjala/datastructures/graph/vertex"
	"golang.org/x/exp/constraints"
)

func bellmanFord[V comparable, W constraints.Ordered](g *graph.Graph[V, W], start, end *vertex.Vertex[V, W]) ([]V, error) {
	pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	delta := make(map[*vertex.Vertex[V, W]]W)

	for _, v := range g.Vertices {
		for edge, w := range v.Edges {
			calcDist := delta[v] + w
			if currDist, ok := delta[edge]; ok {
				if currDist < calcDist {
					continue
				}
			}
			pi[edge] = v
			delta[edge] = calcDist
		}
	}

	for _, v := range g.Vertices {
		for edge, w := range v.Edges {
			calcDist := delta[v] + w
			if currDist, ok := delta[edge]; ok {
				if calcDist < currDist {
					pi[edge] = nil
					delete(delta, edge)
				}
			}
		}
	}
	return path(pi, end), nil
}

func path[V comparable, W constraints.Ordered](
	pi map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W],
	end *vertex.Vertex[V, W]) []V {
	if end == nil {
		return []V{}
	}
	arr := path(pi, pi[end])
	return append(arr, end.GetState())
}
