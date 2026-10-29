package internal

import (
	"bytes"
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

	fileBuffer := make([]byte, 4096)
	var lineBuffer []byte

	file.Seek(-4096, io.SeekEnd)
	file.Read(fileBuffer)

	trimmedBuffer := bytes.TrimRight(fileBuffer, "\x00")

	var arrayIndex int = len(trimmedBuffer) - 1
	var found bool = false

	for !(found) && arrayIndex > 0 {
		if trimmedBuffer[arrayIndex] == 0x0A {
			break
		} else {
			lineBuffer = append(lineBuffer, trimmedBuffer[arrayIndex])
			arrayIndex -= 1
		}
	}

	slices.Reverse(lineBuffer)
	fmt.Printf("%s", lineBuffer)

}
