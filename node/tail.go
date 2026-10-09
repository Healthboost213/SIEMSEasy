package node

import (
	"fmt"
	"io"
	"log"
	"os"
	"slices"
)

func FileTail(fileName string) {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}

	textBuffer := make([]byte, 1)
	var lineData []byte

	var currCount int64
	currCount = 0

	for i := 0; i <= 10; i++ {
		for textBuffer[0] != 0x0A {
			file.Seek(-currCount, io.SeekEnd)
			file.Read(textBuffer)
			currCount += 1
			lineData = append(lineData, textBuffer[0])
		}
		textBuffer[0] = 0x00
		slices.Reverse(lineData)
		fmt.Printf("%x\n", lineData)
		lineData = lineData[:0]
	}

}
