package utils

import "errors"

func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name cannot be empty")
	}
	return "hello " + name, nil

}
