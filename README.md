# errorhandle

Static analysis tool for Go that checks if errors are properly handled

## Install
```
go get -u github.com/Khdbble/errorhandle
```

## what this does

* detect receiving of returning `error` type value with `_` variable
* check if the received `error` type is handled with conditional branch
