package main

import (
	"fmt"
	"os/exec"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOOS)

	cmd := "ls -lah "
	out, err := exec.Command("sh", "-c", cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(out))

	fmt.Println(runCommand("say boobs"))

}

func runCommand(command string) (outString string, errorMessage string) {

	var shell string

	if runtime.GOOS == "windows" {
		shell = "PS"
	} else {
		shell = "sh"
	}

	out, err := exec.Command(shell, "-c", command).Output()
	if err != nil {
		fmt.Println(err.Error())
		errorMessage = err.Error()

		return
	}

	outString = string(out)

	return

}
