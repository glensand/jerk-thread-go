package main

import "jerk-thread-go/containers"

func main() {
	queue := containers.MakeQueue()
	queue.Enqueue("first")
	first := ""
	queue.Dequeue(&first)
	print(first)
}
