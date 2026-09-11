package files

import (
	"os"
	"path/filepath"
	"sort"
)

// Função que retorna os arquivos mais grandes em um diretório
func Largest(root string) ([]Entry, error) {
	var entries []Entry

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}
		entries = append(entries, Entry{
			Path:     path,
			Size:     info.Size(),
			IsDir:    info.IsDir(),
			Modified: info.ModTime(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}
	//Ordena os arquivos por tamanho
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Size > entries[j].Size
	})

	return entries, nil
}
