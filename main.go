package main

import (
	"SIEMSEasy/node"
	"SIEMSEasy/realtime"
)

func main() {

	var queue realtime.RingBuffer
	queue.Head = 0
	queue.Tail = 0

	node.FileTail("README.md")

}
