package main

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
)

func execCmd(s []string) (string, error) {
	if len(s) == 0 {
		return "", errors.New("The command can not be empty!")
	}

	cmd := exec.Command(s[0], s[1:]...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func main() {
	cmdSlice := []string{"runc", "-v"}

	out, err := execCmd(cmdSlice)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
}
