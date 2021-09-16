package fma

import (
	"errors"
	"log"
	"os"
	"os/exec"
)

func CheckFMAConfigToolInstalled() error {

	_, err := exec.LookPath("fma-config-tool")
	if err != nil {
		return errors.New("fma-config-tool is not installed")
	}
	return nil
}

func fmaPath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Could not get the user's home directory: %v", err)
	}
	return dir + "/.local/share/file-manager/actions/"
}

func fmaFilename(name string) string {
	return fmaPath() + name + ".desktop"
}