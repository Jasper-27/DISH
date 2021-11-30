package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"
)

func main() {

	fmt.Println(runtime.GOOS)

	// cmd := "ls -lah "
	// out, err := exec.Command("", "-c", cmd).Output()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Print(string(out))

	fmt.Println(runCommand("say boobs"))

	// Unique IDs

	mac, _ := getMacAddr()

	fmt.Println(strings.Join(mac[:], " "))

}

func runCommand(command string) (outString string, errorMessage string) {

	var shell string

	if runtime.GOOS == "windows" {
		shell = "powershell.exe"
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

func getMacAddr() ([]string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	var as []string
	for _, ifa := range ifas {
		a := ifa.HardwareAddr.String()
		if a != "" {
			as = append(as, a)
		}
	}
	return as, nil
}
