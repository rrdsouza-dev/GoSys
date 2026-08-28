package files

// Importações para o sistema
import (
	"os"
	"path/filepath"
	"strings"
)

// Função de procura que cria slice.
func Search(root, query string) ([]Entry, error) {
	var entries []Entry

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Se for diretório, ignorar
		if info.IsDir() {
			return nil
		}

		// Converter arquivos para letras minúsculas e comparar
		if strings.Contains(
			strings.ToLower(info.Name()),
			strings.ToLower(query),
		) {
			entries = append(entries, Entry{
				Path:  path,
				Size:  info.Size(), // Corrigido: usa dois-pontos ':'
				IsDir: false,
			})
		}
		return nil
	}) // Corrigido: fecha a função anônima e o filepath.Walk

	if err != nil {
		return nil, err
	}

	return entries, nil
}
