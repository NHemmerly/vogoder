package main

import (
	"fmt"
	"os"

	"github.com/go-audio/wav"
)

func main() {
	newFile, _ := os.Create("splicedOut.wav")
	enc := wav.NewEncoder(newFile, 48000, 32, 2, 1)
	splicer := Splicer{soundbank: make(map[string]string), outEncoder: enc}
	splicer.populateSoundbank("soundbank.yml", "./sampleWavs")

	if err := splicer.parseDialogue("textDiag.txt"); err != nil {
		fmt.Printf("error: %s", err)
	}
	enc.Close()
	defer newFile.Close()
	/*
		mp3Path := "testSounds/drum.mp3"
		wavOut := "testSounds/output.wav"
		fmt.Println("Hello World!")
		// Convert .mp3 to .wav
			err := ffmpeg.Input(mp3Path).
				Output(wavOut, ffmpeg.KwArgs{"acodec": "pcm_s16le", "ac": "2", "ar": "48000"}).
				OverWriteOutput().ErrorToStdOut().Run()
			if err != nil {
				log.Fatalf("failed to convert file")
			}
	*/

}
