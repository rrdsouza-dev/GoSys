package files

import (
	"os"
	"path/filepath"
)

type Entry struct {
	Path  string
	Size  int64
	IsDir bool
}

func Scan(root string) ([]Entry, error) {
	var entries []Entry

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		entries = append(entries, Entry{
			Path:  path,
			Size:  info.Size(),
			IsDir: info.IsDir(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}
