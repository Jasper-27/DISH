package main

import (
	"fmt"
	"net"
	"os/exec"
	"runtime"
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

	// mac, _ := getMacAddr()

	// fmt.Println(strings.Join(mac[:], " "))

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

func getMacAddr() (string, error) {
	ifas, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	address := ifas[0].HardwareAddr.String() // gets the MAC address from the first network interface

	return address, nil
}
