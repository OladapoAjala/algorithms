package dijkstra

import (
	"github.com/OladapoAjala/datastructures/graph"
	"github.com/OladapoAjala/datastructures/graph/vertex"
	"github.com/OladapoAjala/datastructures/queues/minpriorityqueue"
	"golang.org/x/exp/constraints"
)

func dijkstra[V comparable, W constraints.Ordered](g *graph.Graph[V, W], start, end *vertex.Vertex[V, W]) ([]V, error) {
	pi := make(map[*vertex.Vertex[V, W]]*vertex.Vertex[V, W])
	delta := make(map[*vertex.Vertex[V, W]]W)
	visited := make(map[*vertex.Vertex[V, W]]bool)
	pq := minpriorityqueue.NewPQueue[W, *vertex.Vertex[V, W]]()

	pi[start] = nil
	delta[start] = *new(W)
	err := pq.Enqueue(*new(W), start)
	if err != nil {
		return nil, err
	}

	for v, err := pq.Dequeue(); err == nil; v, err = pq.Dequeue() {
		visited[v] = true
		for e, w := range v.Edges {
			if visited[e] {
				continue
			}

			calcDist := delta[v] + w
			if currDist, ok := delta[e]; ok {
				if calcDist < currDist {
					pi[e] = v
					delta[e] = calcDist
				}
			} else {
				pi[e] = v
				delta[e] = calcDist
			}
			err := pq.Enqueue(delta[e], e)
			if err != nil {
				return nil, err
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
