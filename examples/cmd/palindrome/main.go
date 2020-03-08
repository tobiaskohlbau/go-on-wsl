package main

import (
	"errors"
	"io"
	"log"
	"os"
)

func main() {
	input := make([]byte, 1024)
	readBytes, err := os.Stdin.Read(input)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}

	output := make([]byte, readBytes)
	for i := readBytes; i > 0; i-- {
		output[readBytes-i] = input[i-1]
	}
	writtenBytes, err := os.Stdout.Write(output)
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}

	if readBytes != writtenBytes {
		log.Fatal("read %d bytes but written only %d\n", readBytes, writtenBytes)
	}
}
