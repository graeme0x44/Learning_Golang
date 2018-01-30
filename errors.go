package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// For no error return nil
	return arg + 3, nil
}

// Custom error type
type argError struct {
	arg int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

// Using custom error type
func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "Cant use this"}
	}
	return arg + 3, nil
}

func main() {
	// Test f1() error
	for _, i := range []int{7, 42}{
		// Inline error check, common Go idiom
		if r, e := f1(i); e != nil {
			fmt.Println("f1 fail", e)
		} else {
			fmt.Println("f1 worked", r)
		}
	}
	// Test f2() error
	for _, i := range []int{7,42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 fail", e)
		} else {
			fmt.Println("f2 worked", r)
		}
	}

	_, e := f2(42)
	// Get custom error type via assertion
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}


