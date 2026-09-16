package system

import (
	"os/exec"
	"strconv"
	"strings"
)

func Memory() (float64, error) {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-CimInstance Win32_OperatingSystem).FreePhysicalMemory",
	)

	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	value := strings.TrimSpace(string(output))
	freeMemory, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return freeMemory, nil
}
