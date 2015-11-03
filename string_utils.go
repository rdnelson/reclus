package main

import "errors"

func ListContains(lst []string, val string) (int, error) {
	for i, b := range lst {
		if b == val {
			return i, nil
		}
	}

	return 0, errors.New("The value was not found in the list")
}
