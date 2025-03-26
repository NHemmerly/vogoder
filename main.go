package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/go-mp3"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	mp3Path := "soundbank/drum.mp3"
	wavOut := "soundbank/output.wav"
	fmt.Println("Hello World!")
	// Sample code from github repo
	fileBytes, err := os.ReadFile(mp3Path)
	if err != nil {
		log.Fatalf("could not read soundFile: %s", err)
	}

	fileBytesReader := bytes.NewReader(fileBytes)

	decodedMp3, err := mp3.NewDecoder(fileBytesReader)
	if err != nil {
		log.Fatalf("mp3.NewDecoder failed: %s", err)
	}
	fmt.Println(decodedMp3.Length())

	err = ffmpeg.Input("")

}
