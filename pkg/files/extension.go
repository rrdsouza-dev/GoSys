package files

import (
	"path/filepath"
	"strings"
)

// Função de pegar a extensão de um arquivo
func GetExtension(path string) string {
	extencion := filepath.Ext(path)
	return strings.ToLower(extencion)
}
