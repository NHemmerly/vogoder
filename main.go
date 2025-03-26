package main

import (
	"fmt"
	"log"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func main() {
	mp3Path := "soundbank/drum.mp3"
	wavOut := "soundbank/output.wav"
	fmt.Println("Hello World!")
	// Convert .mp3 to .wav

	err := ffmpeg.Input(mp3Path).
		Output(wavOut, ffmpeg.KwArgs{"acodec": "pcm_s16le", "ac": "2", "ar": "48000"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		log.Fatalf("failed to convert file")
	}

}
