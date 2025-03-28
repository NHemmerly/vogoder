package main

import (
	"fmt"
	"os"
)

func (s *Splicer) parseDialogue(dialogFile string) error {
	dat, err := os.ReadFile(dialogFile)
	if err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}
	dialogue := string(dat)
	for i := 0; i < len(dialogue); i++ {
		phon := string(dialogue[i])
		fmt.Printf("%s\n", phon)
		if phon == " " {
			phon = s.soundbank["space"]
			rawSound, err := decodeFile(phon)
			if err != nil {
				return fmt.Errorf("error decoding sound file: %w", err)
			}
			if err := s.encodeToOut(rawSound); err != nil {
				return fmt.Errorf("error encoding sound: %w", err)
			}
			fmt.Printf("%s", phon)
			continue
		}
		phon, ok := s.soundbank[string(dialogue[i])]
		fmt.Printf("%v", phon)
		if !ok {
			phon, ok = s.soundbank[string(dialogue[i:i+2])]
			if !ok {
				return fmt.Errorf("phon does not exist or not read properly")
			}
			i++
		}
		fmt.Printf("%sblah", s.soundbank[phon])
		rawSound, err := decodeFile(phon)
		if err != nil {
			return fmt.Errorf("error decoding sound file: %w", err)
		}
		if err := s.encodeToOut(rawSound); err != nil {
			return fmt.Errorf("error encoding sound: %w", err)
		}
		fmt.Printf("%s", phon)

	}
	return nil
}
