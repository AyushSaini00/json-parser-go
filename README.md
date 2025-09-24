trying to build a smol json-parser in go, part of this challenge: https://codingchallenges.fyi/challenges/challenge-json-parser

approach & pseudocode:

- json is made up of objects, arrays, strings, numbers, booleans, and null.
- ex of a valid json string: `{"name": "Rohan", "interested": ["hiking", "coffee"]}`
- steps to parse it:
  - Tokenize: scan the input string into meaningful chunks:
    - `{`, `"name"`, `:`, `"Rohan"`, ...
  - Parsing: Parse these tokens using recursive descent

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
