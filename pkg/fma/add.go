package fma

import (
	"fmt"
	"os"
)

func AddFMA(name, command string, targetsContext, targetsLocation bool) error {
	err := CheckFMAConfigToolInstalled()
	if err != nil {
		return err
	}
	err = os.MkdirAll(fmaPath(), os.ModeDir)
	if err != nil {
		return err
	}
	fileName := fmaFilename(name)
	fileContent := fmt.Sprintf(fmaTemplate, name, command, targetsContext, targetsLocation)
	err = os.WriteFile(fileName, []byte(fileContent), os.FileMode(0777))
	if err != nil {
		return err
	}
	return nil
}

const fmaTemplate string =
`[Desktop Entry]
Type=Action

Name[en_GB]=%[1]v
Name[en]=%[1]v
Name[C]=%[1]v
Profiles=profile-zero;

TargetContext=%[3]v
TargetLocation=%[4]v

[X-Action-Profile profile-zero]
Exec=%[2]v
Name[en_GB]=%[1]v
Name[en]=%[1]v
Name[C]=%[1]v`