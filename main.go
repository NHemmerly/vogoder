package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	//handle command line args
	outputPtr := flag.String("o", "output.wav", "final wav output filepath")
	soundbankPtr := flag.String("s", "soundbank.yml", "yml file that determines the vowels and consonants of the language")
	soundDirPtr := flag.String("d", "soundbank", "filepath of the directory where source sounds reside")
	dialoguePtr := flag.String("t", "dialogue.txt", "filepath to the textfile that will be parsed for audio output")
	sampleRatePtr := flag.Int("sampleRate", 48000, "sample rate to use for the output file")
	bitDepthPtr := flag.Int("bitDepth", 32, "bit depth for the final output file")
	channelsPtr := flag.Int("numChannels", 2, "number of audio channels to use for final output file")
	flag.Parse()
	//initialize encoder
	newFile, _ := os.Create(*outputPtr)
	enc := wav.NewEncoder(
		newFile,
		*sampleRatePtr,
		*bitDepthPtr,
		*channelsPtr,
		1)
	splicer := Splicer{soundbank: make(map[string]string), outEncoder: enc}
	splicer.populateSoundbank(*soundbankPtr, *soundDirPtr)

	if err := splicer.parseDialogue(*dialoguePtr); err != nil {
		fmt.Printf("error: %s", err)
	}
	enc.Close()
	defer newFile.Close()

}
