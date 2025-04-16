package scanner

import (
	"crypto/sha256"
	"fmt"
	"github.com/schollz/progressbar/v3"
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

	bar := progressbar.Default(int64(len(files)), "Hashing files")

	dupeMap := make(map[string][]string)
	for _, file := range files {
		hash, err := hashFile(file)
		if err == nil {
			dupeMap[hash] = append(dupeMap[hash], file)
		}
		err = bar.Add(1)
		if err != nil {
			return nil, err
		}
	}

	return dupeMap, nil
}
