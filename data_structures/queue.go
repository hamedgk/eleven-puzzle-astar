package data_structures

// An Item is something we manage in a priority queue.
//type Item struct {
//	value    string // The value of the item; arbitrary.
//	priority int    // The priority of the item in the queue.
//	// The index is needed by update and is maintained by the heap.Interface methods.
//	index int // The index of the item in the heap.
//}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []Node

func (piq PriorityQueue) Len() int {
	return len(piq)
}
func (piq PriorityQueue) Less(i, j int) bool {

	return piq[i].HeuristicCost < piq[j].HeuristicCost
}
func (piq PriorityQueue) Swap(i, j int) {
	piq[i], piq[j] = piq[j], piq[i]
}
func (piq *PriorityQueue) Push(x any) {
	item := x.(Node)
	*piq = append(*piq, item)
}
func (piq *PriorityQueue) Pop() any {
	old := *piq
	n := len(old)
	item := old[n-1]
	*piq = old[0 : n-1]
	return item
}
