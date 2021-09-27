package alias

import (
	"d/pkg/bashrc"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	filename, err := bashrc.Filename()
	if err != nil {
		return err
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	r, err := regexp.Compile(fmt.Sprintf(`\nalias %v=.*\s*`, regexp.QuoteMeta(s)))
	if err != nil {
		return err
	}
	content = r.ReplaceAll(content, []byte{})
	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}