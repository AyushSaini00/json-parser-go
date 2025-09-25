package main

import (
	"fmt"
	"io"
	"os"

	"github.com/AyushSaini00/json-parser-go/internal/parser"
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

	fileContentBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	if len(string(fileContentBytes)) == 0 {
		fmt.Println("Invalid JSON")
		os.Exit(1)
	}

	err = parseJSON(string(fileContentBytes))
	if err != nil {
		os.Exit(1)
	}
}

func parseJSON(input string) error {
	tokens, err := tokenizer.Tokenize(input)
	if err != nil {
		return err
	}

	// for _, token := range tokens {
	// 	fmt.Printf("--%s: %s\n", token.Type, token.Value)
	// }

	res, err := parser.ParseTokens(tokens)
	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", res)

	return nil
}
