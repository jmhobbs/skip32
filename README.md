[![Build Status](https://circleci.com/gh/jmhobbs/skip32/tree/main.svg?style=shield)](https://circleci.com/gh/jmhobbs/skip32/tree/main)
[![codecov](https://codecov.io/gh/jmhobbs/skip32/branch/main/graph/badge.svg)](https://codecov.io/gh/jmhobbs/skip32)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jmhobbs/skip32)](https://pkg.go.dev/github.com/jmhobbs/skip32)
[![Release](https://img.shields.io/github/release/jmhobbs/skip32.svg?style=flat-square)](https://github.com/jmhobbs/skip32/releases/latest)

# Skip32 for Go

This is a (more or less) direct Golang port of SKIP32 written by Greg Rose

https://web.archive.org/web/20110819120213/http://www.qualcomm.com.au/PublicationsDocs/skip32.c

This cipher is useful for obfuscating 32-bit values in the same size output space.

This can be useful for exposing an incremental ID externally without leaking sequence or value.

## Usage

```go
package main

import (
	"fmt"

	"github.com/jmhobbs/skip32"
)

func main() {
	key := [10]byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09}
	fmt.Println(skip32.Encrypt(key, 500))
	// 499237320

	anotherKey := skip32.KeyFromSlice([]byte("this is a string"))
	fmt.Println(skip32.Encrypt(anotherKey, 500))
	// 1928549585
}
```