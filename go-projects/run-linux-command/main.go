package main

import (
	"fmt"
	"os/exec"
)

func main() {
	listDir, err := exec.Command("ls /Users/wongjunyao/Documents", "-la").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	ls := string(listDir[:])
	fmt.Println(ls)

	printWorkingDirectory, err := exec.Command("pwd").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	pwd := string(printWorkingDirectory[:])
	fmt.Println(pwd)

	portExploit, err := exec.Command("nmap yahoo.com").Output()

	if err != nil {
		fmt.Printf("%s", err)
	}

	nmap := string(portExploit[:])
	fmt.Println(nmap)

}
