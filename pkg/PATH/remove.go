package PATH

import (
	"d/pkg/bashrc"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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
	filename, err := bashrc.Filename()
	if err != nil {
		return err
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	r, err := regexp.Compile(fmt.Sprintf(`\nexport PATH="%v:\$PATH"`, regexp.QuoteMeta(s)))
	if err != nil {
		return err
	}
	if !r.Match(content) {
		return errors.New(fmt.Sprintf("The path %v is in the PATH but not exported via the ~/.bashrc file", s))
	}
	content = r.ReplaceAll(content, []byte{})
	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}