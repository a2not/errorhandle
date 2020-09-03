package a

import "errors"

type myinterface interface{}

type myerr struct{}

func (e *myerr) Error() string {
	return "myerr"
}

func f() (int, error) {
	return 0, nil
}

var (
	_, _    = f() // want "receiving error with _"
	_, err1 = f() // OK
)

// exceptional case pointed out by @110y
// https://github.com/Khdbble/errorhandle/pull/3#issuecomment-686300397
var _ error = (*myerr)(nil) // OK


func main() {
	_, err := f() // OK
	if err != nil {
		panic(err)
	}

	_, _ = f() // want "receiving error with _"

	a, _ := f() // want "receiving error with _"

	a, _ = 12, error(nil) // want "receiving error with _"

	_, _ = 12, "foo" // OK

	m := map[string]error{
		"error1": errors.New("error1"),
	}

	_, ok := m["error1"] // want "receiving error with _"
	if !ok {
	}

	_ = &myerr{} // want "receiving error with _"

	var i myinterface

	_, ok = i.(error) // want "receiving error with _"
	if !ok {
	}

	print(a)
}
