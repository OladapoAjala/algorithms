package dynamicprogramming

import (
	"github.com/OladapoAjala/datastructures/graph/vertex"
	"golang.org/x/exp/constraints"
)

// f(s,v) = min(f(s,u) + w(u,v))

func ShortestPath[V comparable, W constraints.Ordered](s, v *vertex.Vertex[V, W]) []V {
	pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	delta := make(map[*vertex.Vertex[V, W]]W)
	pi[s] = nil
	delta[s] = *new(W)

	shortestPath(s, pi, delta)
	return path(v, pi)
}

func path[V comparable, W constraints.Ordered](end *vertex.Vertex[V, W], pi map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W]) []V {
	if end == nil {
		return []V{}
	}

	o := path(pi[end], pi)
	return append(o, end.GetState())
}

func shortestPath[V comparable, W constraints.Ordered](s *vertex.Vertex[V, W], pi map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W], delta map[*vertex.Vertex[V, W]]W) {
	for e, w := range s.Edges {
		if _, visited := delta[e]; visited {
			currDelta := delta[e]
			calcDelta := delta[s] + w
			if currDelta < calcDelta {
				continue
			}
		}

		pi[e] = s
		delta[e] = delta[s] + w
		shortestPath(e, pi, delta)
	}
}
