package files

import (
	"os"
	"path/filepath"
)

// Função inicial para encontrar arquivos vazios em um diretório
func Empty(root string) ([]Entry, error) {
	var entries []Entry
	//Percorre o diretório e seus subdiretórios
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if info.Size() == 0 {
			entries = append(entries, Entry{
				Path:  path,
				Size:  info.Size(),
				IsDir: false,
			})
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}
