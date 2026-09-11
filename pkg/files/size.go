package files

import (
	"os"
	"path/filepath"
)

// Função para retornar o tamanho de todos os arquivos em um diretório
func Size(root string) ([]Entry, error) {
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
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return entries, nil
}
