# go-generate-password

[![GoDoc](https://godoc.org/github.com/m1/go-generate-password?status.svg)](https://godoc.org/github.com/m1/go-generate-password)
[![Build Status](https://travis-ci.org/m1/go-generate-password.svg?branch=master)](https://travis-ci.org/m1/go-generate-password)
[![Go Report Card](https://goreportcard.com/badge/github.com/m1/go-generate-password)](https://goreportcard.com/report/github.com/m1/go-generate-password)
[![Release](https://img.shields.io/github/release/m1/go-generate-password.svg)](https://github.com/m1/go-generate-password/releases/latest)
[![codecov](https://codecov.io/gh/m1/go-generate-password/branch/master/graph/badge.svg)](https://codecov.io/gh/m1/go-generate-password)

__Password generator written in Go.  Use as a [library](#library) or as a [CLI](#cli).__

## Usage

### CLI

go-generate-password can be used on the cli, just install using: `go get github.com/m1/go-generate-password/cmd/go-generate-password`

To use:
```
➜  go-generate-password --help
go-generate-password is a password generating engine written in Go.

Usage:
  go-generate-password [flags]

Flags:
      --characters string   Character set for the config
      --exclude-ambiguous   Exclude ambiguous characters (default true)
      --exclude-similar     Exclude similar characters (default true)
  -h, --help                help for go-generate-password
  -l, --length int          Length of the password (default 24)
      --lowercase           Include lowercase letters (default true)
      --numbers             Include numbers (default true)
      --symbols             Include symbols (default true)
  -n, --times int           How many passwords to generate (default 1)
      --uppercase           Include uppercase letters (default true)
```

For example: 
```
➜  go-generate-password
5PU?rG-w9YkDus4?AbmKd+Z*
```

More detailed example:
```
➜  go-generate-password -n 5 --length 16 --symbols=false
X89R4HvATgg7HSKk
YMwMPnXp7cnMTNdZ
RZWKAvyxFxDWRB8u
PvKb6uP4N7vAMVsD
KHttvhevGrTYptM5
```

Example using custom character set:
```.env
➜  go-generate-password -n 5 --characters=abcdefg01       
10cecfcfe0bea1fdcbb1afcf
bfcgbgg0dccafdacdaa1de01
gb0ggcffcefae0bb1ac0bbge
abafbc1bbaff0cfbdgaee11d
1fge0fcbccabda0g0a01ffc0
```

### Library
To use as a library is pretty simple:

```go
config := generator.Config{
		Length:                     16,
		IncludeSymbols:             false,
		IncludeNumbers:             true,
		IncludeLowercaseLetters:    true,
		IncludeUppercaseLetters:    true,
		ExcludeSimilarCharacters:   true,
		ExcludeAmbiguousCharacters: true,
	}
g, _ := generator.New(&config)

pwd, _ := g.Generate() 
// pwd = 8hp43B2R7gaXrZUW

pwds, _ := g.GenerateMany(5)
// pwds = [
//   dnPp2TW2e8wmkAwT,
//   XVYwWn25xuNwhUTy,
//   vQ8aSrustQzxQCkA,
//   AuT4fu5RU9TtxEUR,
//   muDwwBRpKpC5BcHr,
// ]

pwd, _ = g.GenerateWithLength(12)
// pwd = HHhpzRGsmEWt

pwds, _ := g.GenerateManyWithLength(5, 12)
// pwds = [
//   s5TKYPdgRzvZ
//   wZFgzs8PUvRg
//   tU73qZ9sPzEs
//   mMaYU6hkvxPQ
//   KBNZ2D7cVQS2
// ]
```

The library also comes with some helper constants:
```go
const (
	// LengthWeak weak length password
	LengthWeak = 6

	// LengthOK ok length password
	LengthOK = 12

	// LengthStrong strong length password
	LengthStrong = 24

	// LengthVeryStrong very strong length password
	LengthVeryStrong = 36

	// DefaultLetterSet is the letter set that is defaulted to - just the
	// alphabet
	DefaultLetterSet = "abcdefghijklmnopqrstuvwxyz"

	// DefaultLetterAmbiguousSet are letters which are removed from the
	// chosen character set if removing similar characters
	DefaultLetterAmbiguousSet = "ijlo"

	// DefaultNumberSet the default symbol set if character set hasn't been
	// selected
	DefaultNumberSet = "0123456789"

	// DefaultNumberAmbiguousSet are the numbers which are removed from the
	// chosen character set if removing similar characters
	DefaultNumberAmbiguousSet = "01"

	// DefaultSymbolSet the default symbol set if character set hasn't been
	// selected
	DefaultSymbolSet = "!$%^&*()_+{}:@[];'#<>?,./|\\-=?"

	// DefaultSymbolAmbiguousSet are the symbols which are removed from the
	// chosen character set if removing ambiguous characters
	DefaultSymbolAmbiguousSet = "<>[](){}:;'/|\\,"
)
```
