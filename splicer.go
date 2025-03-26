package main

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type splicer struct {
	soundbank map[string]string
}

type Soundbank struct {
	Vowels     []string `yaml:"vowels"`
	Consonants []string `yaml:"consonants"`
}

func (s *splicer) populateSoundbank(soundbankPath, dir string) error {
	//Requires a specific file structure to properly populate map
	yamlFile, err := os.ReadFile(soundbankPath)
	if err != nil {
		return fmt.Errorf("error reading file: %s: %s", soundbankPath, err)
	}
	var soundbank Soundbank
	err = yaml.Unmarshal(yamlFile, &soundbank)
	if err != nil {
		return fmt.Errorf("error unmarshalling file: %s", err)
	}
	for _, vowel := range soundbank.Vowels {
		_, ok := s.soundbank[vowel]
		if !ok {
			s.soundbank[vowel] = 
		}
	}
	return nil
}
