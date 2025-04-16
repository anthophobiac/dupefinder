package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func WalkFiles(root string, includeExt, excludeExt []string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(info.Name()))

		if len(includeExt) > 0 && !contains(includeExt, ext) {
			return nil
		}

		if len(excludeExt) > 0 && contains(excludeExt, ext) {
			return nil
		}

		files = append(files, path)
		return nil
	})
	return files, err
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if strings.ToLower(s) == item {
			return true
		}
	}
	return false
}
