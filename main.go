package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-audio/wav"
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

	file, err := os.Open(wavOut)
	if err != nil {
		log.Fatalf("error reading file: %s", err)
		return
	}
	defer file.Close()

	dec := wav.NewDecoder(file)
	newBuf, err := dec.FullPCMBuffer()
	if err != nil {
		log.Fatalf("error creating buffer: %s", err)
		return
	}
	newFile, _ := os.Create("splicedOut.wav")
	enc := wav.NewEncoder(newFile, newBuf.Format.SampleRate, int(dec.BitDepth), newBuf.Format.NumChannels, int(dec.WavAudioFormat))
	enc.Write(newBuf)
	enc.Write(newBuf)
	enc.Write(newBuf)

	enc.Close()
	newFile.Close()

}
