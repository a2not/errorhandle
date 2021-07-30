# errorhandle

[![Go](https://github.com/a2not/errorhandle/actions/workflows/go.yml/badge.svg)](https://github.com/a2not/errorhandle/actions/workflows/go.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

Static analysis tool for Go that checks if errors are assigned to blank identifiers

https://pkg.go.dev/github.com/a2not/errorhandle

## Install
```
go get -u github.com/a2not/errorhandle
```

## what this does

* detect receiving of returning `error` type value with `_` variable

## how to use it

```
make build
export PATH=$PATH:~/go/src/github.com/a2not/errorhandle/cmd/errorhandle
go vet -vettool=`which errorhandle`
```
