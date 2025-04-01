package main

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	//handle command line args
	//initialize encoder
	newFile, _ := os.Create("splicedOut.wav")
	enc := wav.NewEncoder(newFile, 48000, 32, 2, 1)
	splicer := Splicer{soundbank: make(map[string]string), outEncoder: enc}
	splicer.populateSoundbank("soundbank.yml", "./sampleWavs")

	if err := splicer.parseDialogue("textDiag.txt"); err != nil {
		fmt.Printf("error: %s", err)
	}
	enc.Close()
	defer newFile.Close()

}
