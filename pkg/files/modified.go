package files

import (
	"os"
	"path/filepath"
)

func Modified(root string) ([]Entry, error) {
	var entries []Entry

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		entries = append(entries, Entry{
			Path:  path,
			Size:  info.Size(),
			IsDir: false,
			//Importado time(Unica diferença kkk)
			Modified: info.ModTime(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}
