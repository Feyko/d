package alias

import (
	"d/pkg/bashrc"
	"errors"
	"fmt"
	"regexp"
)

func Remove(s string) error {
	exists, err := Exists(s)
	if err != nil {
		return err
	}
	if !exists {
		return errors.New(fmt.Sprintf("The provided alias '%v' doesn't exist", s))
	}
	err = bashrc.ReplaceAll(fmt.Sprintf(`\nalias %v=.*\s*`, regexp.QuoteMeta(s)), "")
	if err != nil {
		return err
	}
	return nil
}