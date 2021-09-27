package alias

import (
	"d/pkg/bashrc"
	"fmt"
	"os"
	"regexp"
)

func Exists(s string) (bool, error) {
	filename, err := bashrc.Filename()
	if err != nil {
		return false, err
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, err
	}
	match, err := regexp.Match(fmt.Sprintf(`\nalias %v=.*\s*`, regexp.QuoteMeta(s)), content)
	if err != nil {
		return false, err
	}
	return match, nil
}