package bashrc

import "os"

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