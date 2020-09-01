# errorhandle

Static analysis tool for Go that checks if errors are properly handled

## Install
```
go get -u github.com/Khdbble/errorhandle
```

## what this does

* detect receiving of returning `error` type value with `_` variable
* check if the received `error` type is handled with conditional branch
* check if the received `error` type is returned 
  * only if it's in the `func` that returns `error`
  * exception where `panic` is appropriate is possible by explicitly stating it with comment (keywords are TBD)
