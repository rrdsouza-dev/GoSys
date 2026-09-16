package process

import (
	"os/exec"
	"strings"
)

func ListWin() ([]Process, error) {
	cmd := exec.Command("tasklist")

	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")

	var processes []Process

	for _, line := range lines {
		fields := strings.Fields(line)

		if len(fields) < 2 {
			continue
		}

		if fields[0] == "Nome" || strings.HasPrefix(fields[0], "=") {
			continue
		}

		processes = append(processes, Process{
			Name: fields[0],
			PID:  fields[1],
		})
	}

	return processes, nil
}
