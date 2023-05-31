package containers

import (
	"sync/atomic"
)

type Node struct {
	data string
	next atomic.Pointer[Node]
}

type Queue struct {
	tail     atomic.Pointer[Node]
	head     *Node
	first    *Node
	tailCopy *Node
}

func MakeQueue() *Queue {
	newNodeRaw := &Node{}
	var newNode atomic.Pointer[Node]
	newNode.Store(newNodeRaw)
	return &Queue{
		tail:     newNode,
		head:     newNodeRaw,
		first:    newNodeRaw,
		tailCopy: newNodeRaw,
	}
}

func (q *Queue) Enqueue(inData string) {
	newNode := &Node{
		data: inData,
	}
	q.head.next.Store(newNode)
	q.head = newNode
}

func (q *Queue) Dequeue(outData *string) bool {
	next := q.tail.Load().next.Load()
	if next != nil {
		*outData = next.data
		q.tail.Store(next)
		return true
	}

	return false
}
