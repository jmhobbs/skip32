[![Build Status](https://circleci.com/gh/jmhobbs/skip32/tree/main.svg?style=shield)](https://circleci.com/gh/jmhobbs/skip32/tree/main)
[![codecov](https://codecov.io/gh/jmhobbs/skip32/branch/main/graph/badge.svg)](https://codecov.io/gh/jmhobbs/skip32)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/jmhobbs/skip32)](https://pkg.go.dev/github.com/jmhobbs/skip32)
[![Release](https://img.shields.io/github/release/jmhobbs/skip32.svg?style=flat-square)](https://github.com/jmhobbs/skip32/releases/latest)

# Skip32 for Go

This is a (more or less) direct Golang port of SKIP32 written by Greg Rose

https://web.archive.org/web/20110819120213/http://www.qualcomm.com.au/PublicationsDocs/skip32.c

This cipher is useful for obfuscating 32-bit values in the same size output space.

This can be useful for exposing an incremental ID externally without leaking sequence or value.
