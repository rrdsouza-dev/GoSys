package files

import (
	"context"
	"io/fs"
	"path/filepath"
)

// Entry representa um arquivo ou diretório encontrado durante o scan.
type Entry struct {
	Path  string
	Size  int64
	IsDir bool
	Err   error
}

// Scan percorre recursivamente o diretório root e retorna todos os itens encontrados.
func Scan(root string) ([]Entry, error) {
	return ScanContext(context.Background(), root)
}

// ScanContext é a versão com suporte a context para cancelamento.
func ScanContext(ctx context.Context, root string) ([]Entry, error) {
	var entries []Entry

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		// Verifica se o contexto foi cancelado
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		// Se deu erro ao ler o item, registra mas CONTINUA
		if err != nil {
			entries = append(entries, Entry{
				Path: path,
				Err:  err,
			})
			// Se for um diretório que falhou, pula ele
			if d != nil && d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		// Pega as informações completas do arquivo
		info, err := d.Info()
		if err != nil {
			entries = append(entries, Entry{
				Path: path,
				Err:  err,
			})
			return nil
		}

		entries = append(entries, Entry{
			Path:  path,
			Size:  info.Size(),
			IsDir: d.IsDir(),
		})

		return nil
	})

	if err != nil {
		return entries, err
	}

	return entries, nil
}
