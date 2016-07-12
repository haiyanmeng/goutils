package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func execCmd(name string, s ...string) (string, error) {
	cmd := exec.Command(name, s...)
	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func main() {
	out, err := execCmd("runc", "-v")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)

	line := out[:strings.Index(out, "\n")]
	fmt.Println(line)
	v := line[strings.LastIndex(line, " ")+1:]
	fmt.Println(v)
	fmt.Println(len(v))

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cwd)
	/*
		cmdSlice = []string{
			"bash",
			"-c",
			`grep -r "\"google.golang.org/grpc\"" /home/hmeng/go/src/github.com/hmeng-19/ocid/Godeps/Godeps.json -A 1 | grep "\"Rev\"" | cut -d: -f2 | tr -d ' "'`,
		}

	*/
}
