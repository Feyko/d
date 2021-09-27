package PATH

import (
	"d/pkg/bashrc"
	"fmt"
)

func Add(s string) error {
	export := fmt.Sprintf("\n" + `export PATH="%v:$PATH"`, s)
	err := bashrc.Append(export)
	if err != nil {
		return err
	}
	return nil
}