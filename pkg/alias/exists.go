package alias

import (
	"d/pkg/bashrc"
	"fmt"
	"regexp"
)

func Exists(s string) (bool, error) {
	match, err := bashrc.Match(fmt.Sprintf(`\nalias %v=.*\s*`, regexp.QuoteMeta(s)))
	if err != nil {
		return false, err
	}
	return match, nil
}