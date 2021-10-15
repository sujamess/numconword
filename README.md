# numconword [![Coverage Status](https://coveralls.io/repos/github/sujamess/numconword/badge.svg?branch=master)](https://coveralls.io/github/99designs/numconword?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/sujamess/numconword)](https://goreportcard.com/report/github.com/sujamess/numconword) [![Go Reference](https://pkg.go.dev/badge/github.com/sujamess/numconword.svg)](https://pkg.go.dev/github.com/sujamess/numconword)

## What is numconword?
[numconword](https://github.com/sujamess/numconword) is a Go library for converting any type of numbers into readable words
- **numconword** prioritizes big number handling — You could convert numbers as big as you want

## Quick start
1. Installation
``` bash
go get -u github.com/sujamess/numconword
```
2. Let's convert!
``` bash
package main

import (
    "fmt"
    
    "github.com/sujamess/numconword"
)

func main() {
    readableWords, err := numconword.ConvertFloat64(123456.78)
    if err != nil {
        // handle error
    }
    
    // one hundred twenty three thousand fifty six point seventy eight
    fmt.Println(readableWords)
}
```

## Contribution
I'll really appreciate contributions! This library has a lot of things to work on (translation, performance improvement, add more test cases, etc.)
