//Credits: bxcodec - github gist: bxcodec/priority_queue.go

package main

type Point struct {
	X, Y   int
	Cost   int
	fScore int

	Index int // The index of the item in the heap.
}

type PriorityQueue []*Point

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest based on cost as the priority
	// The lower the cost, the higher the priority
	return pq[i].fScore < pq[j].fScore
}

// We just implement the pre-defined function in interface of heap.

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	point := old[n-1]
	point.Index = -1
	*pq = old[0 : n-1]
	return point
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	point := x.(*Point)
	point.Index = n
	*pq = append(*pq, point)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}
