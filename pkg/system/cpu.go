package system

import (
	"os/exec"
	"strconv"
	"strings"
)

// Função CPU retorna a porcentagem de uso da CPU.
func CPU() (float64, error) {
	cmd := exec.Command(
		"powershell",
		"-Command",
		"(Get-CimInstance Win32_Processor | Measure-Object -Property LoadPercentage -Average).Average",
	)
	// Executa o comando e captura a saída.
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	// Remove espaços em branco da saída e converte para float64.
	value := strings.TrimSpace(string(output))

	usage, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, err
	}

	return usage, nil
}
