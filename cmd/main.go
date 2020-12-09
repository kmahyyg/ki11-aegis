//go:generate bash -c -- "cd .. ; fileb0x fileb0x.json"

package main

import (
	"fmt"
	"github.com/kmahyyg/ki11-aegis/assets"
	"github.com/kmahyyg/ki11-aegis/config"
	"log"
	"os"
	"os/exec"
	"os/user"
	"net/http"
	"io/ioutil"
	_ "github.com/kmahyyg/ki11-aegis/assets"
	"strings"
	"time"
)

func runShellCmd(command string){
	cmd := exec.Command("/bin/bash", "-c", "--", command)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}

func detectInternet() bool {
	resp, err := http.Get("https://connect.rom.miui.com/generate_204")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 204
}

func main() {
	// detect if root
	currUser, err := user.Current()
	if err == nil {
		if currUser.Username != "root" {
			fmt.Println("Only ROOT can run this program.")
			os.Exit(1)
		}
	} else {
		os.Exit(255)
	}
	fmt.Println("Check permission done.")
	// pre-inst scripts
	for p := 0; p < len(config.PreInst_Scripts); p++ {
		go runShellCmd(config.PreInst_Scripts[p])
	}
	// run apt
	if detectInternet() {
		fmt.Println("Internet detected.")
		runShellCmd(config.Apt_PreStart)
		fmt.Println("Update package database done.")
		// apt install
		pkgStrs := strings.Join(config.Aptpkgs, " ")
		runShellCmd(config.Apt_Inst + " " + pkgStrs)
		fmt.Println("Install packages via apt done.")
	}
	// create temp folder
	parentTmp := os.TempDir()
	tempBinDir, err := ioutil.TempDir(parentTmp, "*-ki11aegis")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempBinDir)
	fmt.Println("Create Temporary Folder: " + tempBinDir)
	// create extract folder
	err = os.Mkdir(tempBinDir + config.Scripts_path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	err = os.Mkdir(tempBinDir + config.Binaries_path, 0755)
	if err != nil {
		log.Fatal(err)
	}
	// fully extracted third-party scripts from binary, without any online requirement
	filesEmbed, err := assets.WalkDirs("", true)
	if err != nil {
		log.Fatal(err)
	}
	// copy from in-memory to temp
	for j := 0; j < len(filesEmbed); j++ {
		currentEmbeddedFile := filesEmbed[j]
		fmt.Println("Extracting : " + currentEmbeddedFile)
		if !strings.Contains(currentEmbeddedFile,"/") {
			continue
		} // might be bug in future, so make sure you are putting all files in subfolder
		fileContent, err := assets.ReadFile(currentEmbeddedFile)
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile(tempBinDir + "/" + currentEmbeddedFile, fileContent, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// extract all
		_ = os.Mkdir("/usr/local/bin", 0755)
		if strings.HasSuffix(currentEmbeddedFile, ".xz"){
			fmt.Println("Uncompressing... " + currentEmbeddedFile)
			runShellCmd(config.Extract_base + tempBinDir + "/" + currentEmbeddedFile)
		} else if strings.HasSuffix(currentEmbeddedFile, ".sh") {  // run all scripts
			fmt.Println("Running..." + currentEmbeddedFile)
			runShellCmd("bash "+ tempBinDir + "/" + currentEmbeddedFile)  // fix permission problem
		}
	}
	// fix permission problem
	runShellCmd("chmod -R +x /usr/local/bin")
	// clean out via dpkg
	for k := 0; k < len(config.Must_postClean_Scripts); k++ {
		fmt.Println("Run PostClean Scripts...")
		runShellCmd(config.Must_postClean_Scripts[k])
	}
	// clean out files
	for q := 0; q < len(config.Must_postClean_Data); q++ {
		fmt.Println("Clear left files...")
		go runShellCmd(config.Postclean_base + config.Must_postClean_Data[q])
	}
	// ban all aliyun ip
	for n := 0; n < len(config.Must_bannedIPs); n++ {
		fmt.Println("Banning Aliyun IP...")
		runShellCmd(config.Iptables_cmdprefix + config.Must_bannedIPs[n] + config.Iptables_cmdsuffix)
	}
	// persistent iptables
	fmt.Println("Persistent iptables...")
	runShellCmd("iptables-save > /etc/iptables.rules")
	// sleep 10 secs for I/O wait
	runShellCmd("sync")
	time.Sleep(10 * time.Second)
	// self-clean
	runShellCmd("sync")
	_ = os.Remove(os.Args[0])
	fmt.Println("Self-clean done. Please reboot.")
}