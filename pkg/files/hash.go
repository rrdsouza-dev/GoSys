package files

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// Hash calcula o checksum SHA-256 de um arquivo no caminho especificado.
func Hash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", fmt.Errorf("erro ao abrir o arquivo: %w", err)
	}
	defer file.Close()

	hash := sha256.New()

	// Transfere os dados do arquivo diretamente para o hasher.
	// io.Copy faz a leitura em buffers, sendo eficiente em memória para arquivos grandes.
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("erro ao ler o arquivo para hash: %w", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}