package PATH

import (
	"d/pkg/bashrc"
	"errors"
	"fmt"
	"regexp"
)

func Remove(s string) error {
	included, err := Includes(s)
	if err != nil {
		return err
	}
	if !included {
		return errors.New(fmt.Sprintf("The provided path '%v' isn't included in the PATH", s))
	}
	exp := fmt.Sprintf(`\nexport PATH="%v:\$PATH"`, regexp.QuoteMeta(s))
	match, err := bashrc.Match(exp)
	if err != nil {
		return err
	}
	if !match {
		return errors.New(fmt.Sprintf("The path %v is in the PATH but not exported via the ~/.bashrc file", s))
	}
	err = bashrc.ReplaceAll(exp, "")
	if err != nil {
		return err
	}
	return nil
}