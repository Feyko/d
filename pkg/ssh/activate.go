package ssh

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"os/user"
)

func Activate(username string) error {
	cmd := exec.Command("echo a")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return err
	}
	output, err := ioutil.ReadAll(&out)
	if err != nil {
		return err
	}
	fmt.Println(string(output))
	if err != nil {
		return err
	}
	foundUser, err := user.Lookup(username)
	if err != nil {
		return err
	}
	sshDir := foundUser.HomeDir + "/.ssh"
	files, err := ioutil.ReadDir(sshDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if !file.IsDir() {
			cmd = exec.Command("ssh-add " + sshDir + "/" + file.Name())
			var out bytes.Buffer
			cmd.Stdout = &out
			err = cmd.Run()
			if err != nil {
				return err
			}
			output, err := ioutil.ReadAll(&out)
			if err != nil {
				return err
			}
			fmt.Println(string(output))
		}
	}
	return nil
}