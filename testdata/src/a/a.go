package a

import (
	"fmt"
	"strconv"
)

var (
	a, _ = strconv.Atoi("tt") // want "receiving error with _"
)

func main() {
	c, _ := strconv.Atoi("rr") // want "receiving error with _"
	fmt.Println(c)
}
