package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, error.New("can't work with 42")
	}
	// For no error return nil
	return arg + 3, nil
}
