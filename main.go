package main

import (
	"SIEMSEasy/realtime"
)

func main() {

	var queue realtime.RingBuffer
	queue.Head = 0
	queue.Tail = 0

}
