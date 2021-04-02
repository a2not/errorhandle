# errorhandle

Static analysis tool for Go that checks if errors are assigned to blank identifiers

## Install
```
go get -u github.com/a2not/errorhandle
```

## what this does

* detect receiving of returning `error` type value with `_` variable

