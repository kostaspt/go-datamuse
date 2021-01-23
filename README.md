# go-datamuse

[![Build Status](https://travis-ci.org/kostaspt/go-datamuse.svg?branch=master)](https://travis-ci.org/kostaspt/go-datamuse)
[![Coverage Status](https://coveralls.io/repos/github/kostaspt/go-datamuse/badge.svg?branch=master)](https://coveralls.io/github/kostaspt/go-datamuse?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/kostaspt/go-datamuse)](https://goreportcard.com/report/github.com/kostaspt/go-datamuse)

A [Go(lang)](https://golang.org/) library for the [Datamuse API](https://www.datamuse.com/api/).

## Installation

```bash
$ go get -u github.com/kostaspt/go-datamuse
```

## Example

```go
package main

import "github.com/kostaspt/go-datamuse"

func main() {
	ml, _ := datamuse.New().Words().MeansLike("ringing in the ears").Get()
	// [{Word:tinnitus Score:51691 SyllablesCount:0 Tags:[syn n]} ...]
}
```

Check [more examples](https://github.com/kostaspt/go-datamuse/blob/master/examples_test.go).