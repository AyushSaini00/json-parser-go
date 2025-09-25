package tokenizer

import (
	"fmt"
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
			// no need to skip the closing quote as main loop is incremented
			continue
		}

		// ----------------
		// checking for numbers
		// ----------------
		// Rules for number:
		// - it can be any number, eg: 42, 59, ...
		// - it can have decimals, eg: 1.4, 3.14, ...
		// - it can be any negative number, eg: -1, -49, -2.493, ....
		// - it can have exponents, eg: 1e10, -2.5E-3, 5e+4
		// ----------------

		// valid number can either start with a digit (0-9), or - sign
		if unicode.IsDigit(r) || r == '-' {
			startIndex := i
			i++ //skip to next char

			// if starting char is a number, then the very next char could be any of: number, decimal, e/E
			// if starting char is a - sign, then the very next char should only be number

			for i < len(str) && unicode.IsDigit(rune(str[i])) {
				i++
			}

			if r == '-' {
				for i < len(str) && unicode.IsDigit(rune(str[i])) {
					i++
				}
			}

			// decimal part
			if i < len(str) && rune(str[i]) == '.' {
				i++
				for i < len(str) && unicode.IsDigit(rune(str[i])) {
					i++
				}
			}

			// exponent part
			if i < len(str) && (str[i] == 'e' || str[i] == 'E') {
				i++
				if i < len(str) && (str[i] == '+' || str[i] == '-') {
					i++
				}
				for i < len(str) && unicode.IsDigit(rune(str[i])) {
					i++
				}
			}

			numVal := str[startIndex:i]
			tokens = append(tokens, Token{"NUMBER", numVal})

			i-- // because loop will also increase the counter, this is to reset it.
			continue
		}

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

			continue
		}

		// if none of the conditions matched, it's an unexpected character
		return []Token{}, fmt.Errorf("unexpected character: '%c' at position %d", r, i)
	}

	return tokens, nil
}
