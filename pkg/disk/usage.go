package disk

import (
	"os"
	"path/filepath"
)

// Função que retorna o uso de disco em bytes
func GetUsage(root string) (int64, error) {
	var total int64
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Adiciona o tamanho do arquivo ao total
		total += info.Size()
		return nil
	})

	if err != nil {
		return 0, err
	}

	return total, nil
}
