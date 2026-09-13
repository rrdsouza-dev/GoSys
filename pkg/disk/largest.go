package disk

import (
	"os"
	"path/filepath"
	"sort"
)

// Estrutura que representa um arquivo
type File struct {
	Path string
	Size int64
}

// Função que retorna os arquivos mais grandes em um diretório
func Largest(root string) ([]File, error) {
	var files []File

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Verifica se é um diretório
		if info.IsDir() {
			return nil
		}
		// Adiciona o arquivo à lista
		files = append(files, File{
			Path: path,
			Size: info.Size(),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}
	// Ordena os arquivos por tamanho
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size > files[j].Size
	})

	return files, nil
}
