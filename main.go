package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/AyushSaini00/json-parser-go/internal/tokenizer"
)

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var file *os.File

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// taking in file from piped cmd or redirected
		file = os.Stdin
	} else {
		args := os.Args
		if len(args) < 2 {
			fmt.Println("Expected a file path")
			os.Exit(1)
		}

		filePath := args[1]
		file, err = os.Open(filePath)
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fileContent := scanner.Text()

	if len(fileContent) == 0 {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}

	fmt.Printf("file contents: %v\n", fileContent)

	err = parseJSON(fileContent)
	if err != nil {
		os.Exit(1)
	}
}

func parseJSON(input string) error {
	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return err
	}

	for _, token := range tokens {
		fmt.Printf("--%s: %s\n", token.Type, token.Value)
	}

	return nil
}
