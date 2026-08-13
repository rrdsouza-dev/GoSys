package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rrdsouza-dev/GoSys/pkg/files"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Uso: gosys files scan <diretório>")
		return
	}

	command := os.Args[1]
	subcommand := os.Args[2]

	if command != "files" || subcommand != "scan" {
		fmt.Println("Comando desconhecido")
		return
	}

	root := "."

	if len(os.Args) >= 4 {
		root = os.Args[3]
	}

	entries, err := files.Scan(root)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		if entry.IsDir {
			fmt.Printf("[DIR]  %s\n", entry.Path)
			continue
		}

		fmt.Printf("[FILE] %s (%d bytes)\n", entry.Path, entry.Size)
	}
}
