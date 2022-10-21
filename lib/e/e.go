package e

import "fmt"

func Wrap(message string, err error) error {
	return fmt.Errorf(message, err)
}

func WrapIfErr(message string, err error) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf(message, err)
}
