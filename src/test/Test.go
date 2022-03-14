package main

import (
	"fmt"
	queue2 "go_data_structure/src/queue"
)

func main() {
	var queue *queue2.ItemQueue
	queue = queue2.NewQueue()
	queue.Offer(123)
	queue.Offer(456)
	fmt.Printf("%+v\n", queue.Poll())
	fmt.Printf("%+v", queue.Poll())
}
