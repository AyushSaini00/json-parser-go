package parser

import (
	"fmt"

	"github.com/AyushSaini00/json-parser-go/internal/tokenizer"
)

func ParseTokens(tokens []tokenizer.Token) (map[string]any, error) {
	position := 0
	obj := make(map[string]any)

	for _, token := range tokens {
		if token.Type == "SYMBOL" && token.Value == "{" {
			//start of an object

			position++

			// empty object
			if tokens[position].Type == "SYMBOL" && token.Value == "}" {
				position++
				return obj, nil
			}

			//if not empty object, key must be present
			if tokens[position].Type != "STRING" {
				return nil, fmt.Errorf("expected string key")
			}

			key := tokens[position].Value
			position++

			//expect a colon
			if tokens[position].Value != ":" {
				return nil, fmt.Errorf("expected a colon")
			}
			position++

			//parse the value recursively
			val, err := parseValue(tokens, &position)
			if err != nil {
				return nil, err
			}

			obj[key] = val
			position++

			// next token should either be "," or "}"
			if tokens[position].Value == "," {
				position++
				continue
			} else if tokens[position].Value == "}" {
				position++
				break
			} else {
				return nil, fmt.Errorf("expected , or }")
			}
		}
	}

	return obj, nil
}

func parseValue(tokens []tokenizer.Token, position *int) (interface{}, error) {

	token := tokens[*position]

	switch token.Type {
	case "STRING":
		val := token.Value
		return val, nil
	case "NUMBER":
		val := token.Value
		return val, nil
	case "BOOLEAN":
		val := token.Value == "true"
		return val, nil
	case "NULL":
		return nil, nil
	case "SYMBOL":
		if token.Value == "{" {
			// return parseObject()
		}
		if token.Value == "[" {
			// return parseArray()
		}
		return nil, fmt.Errorf("unexpected symbol: %+v", token)
	default:
		return nil, fmt.Errorf("unexpected token in value: %+v", token)
	}
}
