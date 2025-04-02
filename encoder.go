package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func decodeFile(phonFile string) (*audio.IntBuffer, error) {
	file, err := os.Open(phonFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	defer file.Close()

	dec := wav.NewDecoder(file)
	raw, err := dec.FullPCMBuffer()
	if err != nil {
		return nil, fmt.Errorf("error getting PCM: %w", err)
	}
	return raw, nil
}

func (s *Splicer) encodeToOut(raw *audio.IntBuffer) error {
	if err := s.outEncoder.Write(raw); err != nil {
		return fmt.Errorf("error writing PCM buffer: %w", err)
	}
	return nil
}

func mp3ToWav(mp3Path, outputPath string) error {
	// Convert .mp3 to .wav
	err := ffmpeg.Input(mp3Path).
		Output(outputPath, ffmpeg.KwArgs{"acodec": "pcm_s16le", "ac": "2", "ar": "48000"}).
		OverWriteOutput().ErrorToStdOut().Run()
	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}
	return nil
}
