package system

import (
	"os/exec"
	"strings"
)

func Info() (string, error) {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"Get-CimInstance Win32_OperatingSystem | Select-Object Caption, Version, OSArchitecture",
	)

	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
