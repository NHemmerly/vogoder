package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/go-audio/wav"
	"gopkg.in/yaml.v3"
)

type Splicer struct {
	soundbank  map[string]string
	outEncoder *wav.Encoder
}

type Soundbank struct {
	Vowels     []string `yaml:"vowels"`
	Consonants []string `yaml:"consonants"`
}

func (s *Splicer) checkDirToMap(soundCat []string, dir string) error {
	category, _ := dirExists(dir)
	if !category {
		os.Mkdir(dir, 0755)
	}
	os.Mkdir(dir, 0755)
	for _, phon := range soundCat {
		var phonFile string
		soundDir := dir + phon
		phonExists, err := dirExists(soundDir)
		if err != nil {
			return fmt.Errorf("error checking dir: %w", err)
		}
		if phonExists {
			file, err := os.ReadDir(soundDir)
			if err != nil {
				return fmt.Errorf("error reading dir: %w", err)
			}
			if len(file) > 0 {
				phonFile = file[0].Name()
				s.soundbank[phon] = soundDir + "/" + phonFile
			}
		} else {
			if err := os.Mkdir(soundDir, 0755); err != nil {
				return fmt.Errorf("error making dir: %w", err)
			}
			continue
		}
	}
	return nil
}

func (s *Splicer) populateSoundbank(soundbankPath, dir string) error {
	//Requires a specific file structure to properly populate map
	yamlFile, err := os.ReadFile(soundbankPath)
	if err != nil {
		return fmt.Errorf("error reading file: %s: %s", soundbankPath, err)
	}

	vowelDir := fmt.Sprintf("%s/vowels/", dir)
	consonantDir := fmt.Sprintf("%s/consonants/", dir)
	var soundbank Soundbank
	err = yaml.Unmarshal(yamlFile, &soundbank)
	if err != nil {
		return fmt.Errorf("error unmarshalling file: %s", err)
	}
	if err := s.checkDirToMap(soundbank.Vowels, vowelDir); err != nil {
		return fmt.Errorf("error mapping to dirs: %w", err)
	}
	if err := s.checkDirToMap(soundbank.Consonants, consonantDir); err != nil {
		return fmt.Errorf("error mapping to dirs: %w", err)
	}

	return nil
}

func dirExists(dir string) (bool, error) {
	_, err := os.Stat(dir)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}
