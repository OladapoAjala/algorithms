package dungeon2d

import (
	"fmt"

	"github.com/OladapoAjala/datastructures/queues/queue"
)

func dungeon(grid [][]string, sr, sc int) [][]int {
	R := len(grid)
	C := len(grid[0])
	dr := []int{-1, 1, 0, 0}
	dc := []int{0, 0, 1, -1}

	pi := make(map[int]int)
	rq := queue.NewQueue[int]()
	cq := queue.NewQueue[int]()

	startNode := (sr * C) + sc
	endNode := 0
	foundEnd := false

	_ = rq.Enqueue(sr)
	_ = cq.Enqueue(sc)
	for r, err := rq.Dequeue(); err == nil; r, err = rq.Dequeue() {
		c, err := cq.Dequeue()
		if err != nil {
			return nil
		}
		if grid[r][c] == "#" {
			continue
		}

		currentNode := (r * C) + c
		if grid[r][c] == "s" {
			startNode = currentNode
		}

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
			pi[neighbour] = currentNode
			if grid[rr][cc] == "e" {
				endNode = neighbour
				foundEnd = true
				break
			}
		}
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
