package files

import (
	"os"
)

// função que compara dois arquivos
func Compare(path1, path2 string) (bool, error) {
	file1, err := os.Open(path1)
	if err != nil {
		return false, err
	}
	defer file1.Close()

	file2, err := os.Open(path2)
	if err != nil {
		return false, err
	}
	defer file2.Close()

	info1, err := file1.Stat()
	if err != nil {
		return false, err
	}

	info2, err := file2.Stat()
	if err != nil {
		return false, err
	}

	if info1.Size() != info2.Size() {
		return false, nil
	}

	hash1, err := Hash(path1)
	if err != nil {
		return false, err
	}

	hash2, err := Hash(path2)
	if err != nil {
		return false, err
	}

	return hash1 == hash2, nil
}
