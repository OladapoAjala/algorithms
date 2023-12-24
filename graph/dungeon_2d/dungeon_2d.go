package dungeon2d

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/queues/queue"
)

func dungeon(grid [][]string, sr, sc int) [][]int {
	pi := make(map[int]int)
	rq := queue.NewQueue[int]()
	cq := queue.NewQueue[int]()

	startNode := (sr * len(grid[0])) + sc
	endNode := startNode
	pi[startNode] = -1
	foundEnd := false

	_ = rq.Enqueue(sr)
	_ = cq.Enqueue(sc)
	for r, err := rq.Dequeue(); err == nil; r, err = rq.Dequeue() {
		c, err := cq.Dequeue()
		if err != nil {
			return nil
		}
		endNode, foundEnd = exploreNeighbours(grid, pi, r, c, rq, cq)
		if foundEnd {
			break
		}
	}

	if !foundEnd {
		return nil
	}
	for endNode != startNode {
		fmt.Printf("%d -> ", endNode)
		endNode = pi[endNode]
	}
	fmt.Printf("%d\n", endNode)
	return nil
}

func exploreNeighbours(grid [][]string, pi map[int]int, r int, c int, rq, cq *queue.Queue[int]) (int, bool) {
	R := len(grid)
	C := len(grid[0])
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, 1, -1}

	for i := 0; i < 4; i++ {
		rr := r + dr[i]
		cc := c + dc[i]
		neighbour := (rr * C) + cc

		if _, visited := pi[neighbour]; visited {
			continue
		}
		if rr < 0 || cc < 0 {
			continue
		}
		if rr >= R || cc >= C {
			continue
		}
		if grid[rr][cc] == "#" {
			continue
		}

		_ = rq.Enqueue(rr)
		_ = cq.Enqueue(cc)
		currentNode := (r * C) + c
		pi[neighbour] = currentNode
		if grid[rr][cc] == "e" {
			return neighbour, true
		}
	}
	return -1, false
}
