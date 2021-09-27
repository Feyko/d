package alias

import (
	"d/pkg/bashrc"
	"errors"
	"fmt"
)

func Add(name, exec string) error {
	exists, err := Exists(name)
	if err != nil {
		return errors.New(fmt.Sprintf("Could not check if the alias exists: %v", err))
	}
	if exists {
		return errors.New(fmt.Sprintf("The alias %v already exists", err))
	}
	export := fmt.Sprintf("\n" + `alias %v='%v'`, name, exec)
	err = bashrc.Append(export)
	if err != nil {
		return err
	}
	return nil
}