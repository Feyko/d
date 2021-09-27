package PATH

import (
	"os"
	"strings"
)

func List() ([]string, error) {
	path := os.Getenv("PATH")
	return strings.Split(path, ":"), nil
}