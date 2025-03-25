package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
)

func main() {
	fmt.Println("Hello World!")
	// Sample code from github repo
	fileBytes, err := os.ReadFile("soundbank/drum.mp3")
	if err != nil {
		log.Fatalf("could not read soundFile: %s", err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		log.Fatalf("mp3.NewDecoder failed: %s", err)
	}
	fmt.Println(decodedMp3.Length())
}
