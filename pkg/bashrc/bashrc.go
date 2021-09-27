package bashrc

import (
	"io/ioutil"
	"os"
	"regexp"
)

func Append(t string) error {
	filename, err := Filename()
	if err != nil {
		return err
	}
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	defer file.Close()
	_, err = file.WriteString(t)
	return err
}

func Filename() (string, error) {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homedir + "/.bashrc", nil
}

func Match(exp string) (bool, error) {
	filename, err := Filename()
	if err != nil {
		return false, err
	}
	content, err := os.ReadFile(filename)
	match, err := regexp.Match(exp, content)
	if err != nil {
		return false, err
	}
	return match, nil
}

func ReplaceAll(exp, rep string) error {
	filename, err := Filename()
	if err != nil {
		return err
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	r, err := regexp.Compile(exp)
	if err != nil {
		return err
	}
	content = r.ReplaceAll(content, []byte(rep))
	err = ioutil.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}
	return nil
}