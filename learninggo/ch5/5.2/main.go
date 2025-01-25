package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileLength, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileLength)
}

func fileLen(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	fileBytes := make([]byte, 2028)
	var byteCount int
	for {
		numBytesRead, err := file.Read(fileBytes)
		if err != nil {
			if err == io.EOF {
				break
			}
			return -1, err
		}
		byteCount += numBytesRead
	}
	return byteCount, nil
}
