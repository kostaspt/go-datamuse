# go-datamuse

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/kostaspt/go-datamuse/CI)](https://github.com/kostaspt/go-datamuse/actions)
[![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/kostaspt/go-datamuse)](https://github.com/kostaspt/go-datamuse/releases)
[![GitHub](https://img.shields.io/github/license/kostaspt/go-datamuse)](https://github.com/kostaspt/go-datamuse/blob/main/LICENCE)
[![Go Reference](https://pkg.go.dev/badge/github.com/kostaspt/go-datamuse.svg)](https://pkg.go.dev/github.com/kostaspt/go-datamuse)
[![Coveralls github](https://img.shields.io/coveralls/github/kostaspt/go-datamuse)](https://coveralls.io/github/kostaspt/go-datamuse?branch=master)
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