package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rrdsouza-dev/GoSys/pkg/files"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}

	command := os.Args[1]
	subcommand := os.Args[2]

	switch command {
	case "files":
		handleFilesCommand(subcommand)
	default:
		fmt.Printf("Comando desconhecido: %s\n\n", command)
		usage()
	}
}

func handleFilesCommand(subcommand string) {
	switch subcommand {
	case "scan":
		scanFiles()

	case "search":
		searchFiles()

	case "hash":
		hashFile()

	case "extension":
		extensionFile()

	case "empty":
		if len(os.Args) < 4 {
			fmt.Println("Uso: gosys files empty <diretório>")
			return
		}

		root := os.Args[3]

		entries, err := files.Empty(root)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			fmt.Printf("[EMPTY] %s\n", entry.Path)
		}

	case "size":
		if len(os.Args) < 4 {
			fmt.Println("Uso: gosys files size <diretório>")
			return
		}

		root := os.Args[3]

		entries, err := files.Size(root)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			fmt.Printf("[FILE] %s (%d bytes)\n", entry.Path, entry.Size)
		}

	case "modified":
		if len(os.Args) < 4 {
			fmt.Println("Uso: gosys files modified <diretório>")
			return
		}

		root := os.Args[3]

		entries, err := files.Modified(root)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			fmt.Printf(
				"[FILE] %s — %s\n",
				entry.Path,
				entry.Modified.Format("02/01/2006 15:04:05"),
			)
		}

	case "largest":
		if len(os.Args) < 4 {
			fmt.Println("Uso: gosys files largest <diretório>")
			return
		}

		root := os.Args[3]

		entries, err := files.Largest(root)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			fmt.Printf("[FILE] %s (%d bytes)\n", entry.Path, entry.Size)
		}

	default:
		fmt.Printf("Subcomando desconhecido: %s\n\n", subcommand)
		usage()
	}
}

func scanFiles() {
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

func searchFiles() {
	if len(os.Args) < 5 {
		fmt.Println("Uso: gosys files search <diretório> <nome>")
		return
	}

	root := os.Args[3]
	query := os.Args[4]

	entries, err := files.Search(root, query)
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		fmt.Printf("[FILE] %s (%d bytes)\n", entry.Path, entry.Size)
	}
}

func hashFile() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: gosys files hash <arquivo>")
		return
	}

	path := os.Args[3]

	hash, err := files.Hash(path)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(hash)
}

func extensionFile() {
	if len(os.Args) < 4 {
		fmt.Println("Uso: gosys files extension <arquivo>")
		return
	}

	path := os.Args[3]

	extension := files.GetExtension(path)

	fmt.Println(extension)
}

func usage() {
	fmt.Println("GoSys - System Toolkit")
	fmt.Println()
	fmt.Println("Uso:")
	fmt.Println("  gosys files scan <diretório>")
	fmt.Println("  gosys files search <diretório> <nome>")
	fmt.Println("  gosys files hash <arquivo>")
	fmt.Println("  gosys files extension <arquivo>")
	fmt.Println("  gosys files empty <diretório>")
	fmt.Println("  gosys files size <diretório>")
	fmt.Println("  gosys files modified <diretório>")
	fmt.Println("  gosys files largest <diretório>")
}
