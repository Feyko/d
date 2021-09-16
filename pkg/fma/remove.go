package fma

import (
	"errors"
	"os"
)

func RemoveFMA(name string) error {
	err := os.Remove(fmaFilename(name))
	if os.IsNotExist(err) {
		return errors.New("This FMA doesn't exist!")
	}
	if err != nil {
		return err
	}
	return nil
}
