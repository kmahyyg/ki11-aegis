//go:generate bash -c -- "cd .. ; fileb0x fileb0x.json"

package main

import (
	"fmt"
	"github.com/kmahyyg/ki11-aegis/config"
	"os"
	"os/exec"
	"os/user"
)

func runShellCmd(command string){
	cmd := exec.Command("/bin/bash", "-c", "--", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func main() {
	// detect if root
	currUser, err := user.Current()
	if err == nil {
		if currUser.Username != "root"{
			fmt.Println("Only Root.")
			os.Exit(1)
		}
	} else {
		os.Exit(255)
	}
	runShellCmd(config.Apt_PreStart)
	// apt install
	for i:= 0; i < len(config.Aptpkgs); i++ {
		runShellCmd(config.Apt_Inst + config.Aptpkgs[i])
	}
	// run third-party scripts
	if config.Detect2UseLocal() {

	}
}