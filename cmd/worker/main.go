package main

import (
	"SIEMSEasy/internal/realtime"
)

func main() {

	var queue realtime.RingBuffer
	queue.Head = 0
	queue.Tail = 0

}
