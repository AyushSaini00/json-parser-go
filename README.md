trying to build a smol json-parser in go, part of this challenge: https://codingchallenges.fyi/challenges/challenge-json-parser

approach & pseudocode:

- json is made up of objects, arrays, strings, numbers, booleans, and null.
- ex of a valid json string: `{"name": "Rohan", "interested": ["hiking", "coffee"]}`
- steps to parse:

  - Tokenize: scan the input string into meaningful tokens:
    - `{`
    - `"name"`
    - `:`
    - `"Rohan"`
    - `,`
    - `"interested"`
    - `:`
    - `[`
    - `"hiking"`
    - `,`
    - `"coffee"`
    - `]`
    - `}`
  - Parsing: Parse these tokens using recursive descent

    1. found `{` -> parse the object til we find the closing tag.

    - if not found, it's not a valid json.

    2. found a key `"name"` -> next token should be a colon

    - if not, it's not a valid json.

    3. found a colon `:` -> next token should be any one of the following values:

    - a string
    - a number
    - a boolean
    - null
    - an array opening bracket
    - an object's opening bracket

    - if it's not any of those, then it's an invalid json.

    4. found a string `"Rohan"` -> next token should either be a `,` or `}`

    - if it is a comma, go to next token
    - if it is a closing paran, json is completed and it's a vaild json.

    5. repeating step 2 & 3

    6. found `[` -> parse the array till it's closing tag

    - if not valid array or no closing `]` found, then its an invalid json.

    7. found `]`
    8. found `}` -> json string completed -> it's a valid json

build:

```
go build
```

run:

```
./json-parser-go test.json
```

```
cat test.json | ./json-parser-go
```
