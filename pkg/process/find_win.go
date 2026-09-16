package process

import "strings"

func FindWin(query string) ([]Process, error) {
	processes, err := ListWin()
	if err != nil {
		return nil, err
	}

	var matches []Process

	for _, proc := range processes {
		if strings.Contains(
			strings.ToLower(proc.Name),
			strings.ToLower(query),
		) {
			matches = append(matches, proc)
		}
	}

	return matches, nil
}
