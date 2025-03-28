package main

import (
	"fmt"
	"os"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
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
