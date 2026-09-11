package files

import (
	"os"
	"path/filepath"
)

//Função para encontrar arquivos duplicados em um diretório
func Duplicates(root string) (map[string][]Entry, error) {
	hashes := make(map[string][]Entry)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		hash, err := Hash(path)
		if err != nil {
			return err
		}

		hashes[hash] = append(hashes[hash], Entry{
			Path:     path,
			Size:     info.Size(),
			IsDir:    false,
			Modified: info.ModTime(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return hashes, nil
}
