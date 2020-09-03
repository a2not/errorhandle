package a

func f() (int, error) {
	return 0, nil
}

var (
	_, _ = f() // want "receiving error with _"
	_, err1 = f() // OK
)

func main() {
	_, err := f() // OK
	if err != nil {
		panic(err)
	}

	_, _ = f() // want "receiving error with _"

	a, _ := f() // want "receiving error with _"

	a, _ = 12, error(nil) // want "receiving error with _"

	_, _ = 12, "foo" // OK

	print(a)
}
