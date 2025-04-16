package scanner

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func hashFile(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

func FindDuplicates(root string) (map[string][]string, error) {
	files, err := WalkFiles(root)
	if err != nil {
		return nil, err
	}

	dupeMap := make(map[string][]string)
	for _, file := range files {
		hash, err := hashFile(file)
		if err != nil {
			continue
		}
		dupeMap[hash] = append(dupeMap[hash], file)
	}

	return dupeMap, nil
}
