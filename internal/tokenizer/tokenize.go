package tokenizer

import (
	"slices"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

func Tokenize(str string) ([]Token, error) {
	var tokens []Token
	symbols := []string{
		"{", "}", "[", "]", ":", ",",
	}

	for i := 0; i < len(str); i++ {
		r := rune(str[i])
		// fmt.Printf("--char: %v\n", r)

		// space check
		if unicode.IsSpace(r) {
			continue
		}

		// symbol check
		if slices.Contains(symbols, string(r)) {
			tokens = append(tokens, Token{"SYMBOL", string(r)})
			continue
		}

		// string check
		if r == '"' {
			i += 1 //skip the opening quote
			startIndex := i

			for i < len(str) && rune(str[i]) != '"' {
				i += 1
			}

			strVal := str[startIndex:i]
			tokens = append(tokens, Token{"STRING", strVal})
			i += 1 //skip the closing quote

			continue
		}

		// checking for number
		// an number can either start with '-' or a digit
		// if i < len(str) && (r == '-' || unicode.IsDigit(r)) {
		// 	startIndex := i
		// 	i++ // we can move to the next rune

		// 	for i < len(str) && unicode.IsDigit(rune(str[i])) {
		// 		i++ //move to next run as this is a digit
		// 	}

		// 	for i < len(str) && str[i] == "." {
		// 		i++
		// 	}
		// }

		// bool check
		// check for true
		if i+3 < len(str) && str[i:i+4] == "true" {
			// because we would have the type of BOOLEAN associated with bool, i can store boolean as string
			tokens = append(tokens, Token{"BOOLEAN", "true"})
			i += 3 //skip the rest of true

			continue
		}
		// check for false
		if i+4 < len(str) && str[i:i+5] == "false" {
			tokens = append(tokens, Token{"BOOLEAN", "false"})
			i += 4 //skip the rest of false

			continue
		}

		// checking for null
		if i+3 < len(str) && str[i:i+4] == "null" {
			tokens = append(tokens, Token{"NULL", "null"})
			i += 3 //skip the rest of null
		}

	}

	return tokens, nil
}
