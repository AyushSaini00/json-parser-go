package parser

import (
	"fmt"

	"github.com/AyushSaini00/json-parser-go/internal/tokenizer"
)

func ParseTokens(tokens []tokenizer.Token) (map[string]any, error) {
	position := 0
	return parseObject(tokens, &position)
}

func parseObject(tokens []tokenizer.Token, position *int) (map[string]any, error) {
	obj := make(map[string]any)

	if tokens[*position].Value != tokenizer.OPEN_CURLY_BRACKET {
		return nil, fmt.Errorf("expected {")
	}
	*position++

	// empty object
	if tokens[*position].Value == tokenizer.CLOSE_CURLY_BRACKET {
		*position++
		return obj, nil
	}

	for {

		if tokens[*position].Value == tokenizer.CLOSE_CURLY_BRACKET {
			prev := *position - 1
			if tokens[prev].Value == tokenizer.COMMA {
				return nil, fmt.Errorf("trailing comma")
			}
			break
		}

		//if not empty object, key must be present
		if tokens[*position].Type != tokenizer.STRING {
			return nil, fmt.Errorf("expected string key at %v", *position)
		}

		key := tokens[*position].Value
		*position++

		//expect a colon
		if tokens[*position].Value != tokenizer.COLON {
			return nil, fmt.Errorf("expected a colon")
		}
		*position++

		//parse the value recursively
		val, err := parseValue(tokens, position)
		if err != nil {
			return nil, err
		}

		obj[key] = val

		// next token should either be "," or "}"
		if tokens[*position].Value == tokenizer.COMMA {
			*position++
			continue
		} else if tokens[*position].Value == tokenizer.CLOSE_CURLY_BRACKET {
			*position++
			break
		} else {
			return nil, fmt.Errorf("expected , or } position: %v", *position)
		}

	}

	return obj, nil
}

func parseArray(tokens []tokenizer.Token, position *int) ([]any, error) {
	arr := []any{}

	if tokens[*position].Value != tokenizer.OPEN_SQUARE_BRACKET {
		return nil, fmt.Errorf("expected [")
	}
	*position++

	// empty array
	if tokens[*position].Value == tokenizer.CLOSE_SQUARE_BRACKET {
		*position++
		return arr, nil
	}

	for {
		if tokens[*position].Value == tokenizer.CLOSE_SQUARE_BRACKET {
			prev := *position - 1
			if tokens[prev].Value == tokenizer.COMMA {
				return nil, fmt.Errorf("trailing comma")
			}
			break
		}

		val, err := parseValue(tokens, position)
		if err != nil {
			return nil, err
		}
		arr = append(arr, val)

		if tokens[*position].Value == tokenizer.COMMA {
			*position++
			continue
		} else if tokens[*position].Value == tokenizer.CLOSE_SQUARE_BRACKET {
			*position++
			break
		} else {
			return nil, fmt.Errorf("expected , or ]")
		}
	}

	return arr, nil
}

func parseValue(tokens []tokenizer.Token, position *int) (interface{}, error) {

	token := tokens[*position]

	switch token.Type {
	case tokenizer.STRING:
		val := token.Value
		*position++
		return val, nil
	case tokenizer.NUMBER:
		val := token.Value
		*position++
		return val, nil
	case tokenizer.BOOLEAN:
		val := token.Value == "true"
		*position++
		return val, nil
	case tokenizer.NULL:
		*position++
		return nil, nil
	case tokenizer.SYMBOL:
		if token.Value == tokenizer.OPEN_CURLY_BRACKET {
			return parseObject(tokens, position)
		}
		if token.Value == tokenizer.OPEN_SQUARE_BRACKET {
			return parseArray(tokens, position)
		}
		return nil, fmt.Errorf("unexpected symbol: %+v", token)
	default:
		return nil, fmt.Errorf("unexpected token in value: %+v", token)
	}
}
