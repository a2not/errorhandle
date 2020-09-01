package a

import (
	"fmt"
	"strconv"
)

var (
	_, _ = strconv.Atoi("foo") // want "receiving error with _"
)

func main() {
	_, err := strconv.Atoi("bar") // OK
	if err != nil {
		panic(err)
	}

	_, _ = strconv.Atoi("baz") // want "receiving error with _"

	a, _ := strconv.Atoi("qux") // want "receiving error with _"
	fmt.Println(a)
}
