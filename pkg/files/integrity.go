package files

func Integrity(path, expectedHash string) (bool, error) {
	hash, err := Hash(path)
	if err != nil {
		return false, err
	}

	if hash != expectedHash {
		return false, nil
	}

	return true, nil
}
