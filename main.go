package main

import (
	"container/heap"
	ds "eleven-puzzle/data_structures"
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
)

var examplePuzzle = puzzle.PuzzleBuffer{
	//karimi
	{6, 4, 255, 8},
	{5, 11, 1, 2},
	{7, 9, 3, 10},

	//kazemini example
	//{2, 7, 255, 6},
	//{1, 7, 3, 8},
	//{4, 5, 2, 4},

	//{1, 1, 1, 1},
	//{1, 0, 1, 1},
	//{255, 1, 1, 1},
}

func main() {

	sortedArray := puzzle.SortPuzzle(examplePuzzle)
	digested := puzzle.FromBuffer(sortedArray).Digest()

	explored := map[puzzle.PuzzleBuffer]bool{}
	frontier := make(ds.PriorityQueue, 0)
	heap.Push(
		&frontier,
		ds.Node{
			Parent:        nil,
			Direction:     puzzle.None,
			Cost:          0,
			HeuristicCost: puzzle.Heuristic(digested, examplePuzzle),
			Puzzle:        puzzle.FromBuffer(examplePuzzle),
		},
	)

	for {

		if node, ok := heap.Pop(&frontier).(ds.Node); ok {
			if node.Parent == nil && node.Direction == 0 {
				continue
			}
			if node.IsGoal(sortedArray) {
				ds.TraceBack(&node, 0)
				fmt.Printf("\ntotal nodes are: %d\n", len(explored))
				return
			}
			node.Expand(&frontier, explored, digested)
		}
	}
}
