package data_structures

import (
	"container/heap"
	"eleven-puzzle/data_structures/puzzle"
	"fmt"
)

type Node struct {
	Parent        *Node
	Direction     puzzle.Direction
	Puzzle        puzzle.Puzzle
	Cost          int
	HeuristicCost int
}

func (node *Node) Expand(queue *PriorityQueue, explored map[puzzle.PuzzleBuffer]bool, digested puzzle.DigestCoords) {
	possibleMoves := node.Puzzle.PossibleBlankMoves()
	for _, direction := range possibleMoves {
		copyNode := *node
		copyNode.Direction = direction
		copyNode.Parent = node
		copyNode.Puzzle.MoveBlank(direction)
		copyNode.Cost = node.Cost + 1
		copyNode.HeuristicCost = copyNode.Cost + puzzle.Heuristic(digested, copyNode.Puzzle.Buffer)
		if explored[copyNode.Puzzle.Buffer] {
			continue
		}
		heap.Push(queue, copyNode)
		explored[copyNode.Puzzle.Buffer] = true
	}
}

func (node *Node) IsGoal(buffer puzzle.PuzzleBuffer) bool {
	return node.Puzzle.Buffer == buffer
}

func TraceBack(node Node) {
	if node.Parent == nil {
		fmt.Print("init")
		return
	}

	TraceBack(*node.Parent)
	switch node.Direction {
	case puzzle.Up:
		fmt.Print(" -> Up")
	case puzzle.Down:
		fmt.Print(" -> Down")
	case puzzle.Right:
		fmt.Print(" -> Right")
	case puzzle.Left:
		fmt.Print(" -> Left")
	default:
		return
	}
}
