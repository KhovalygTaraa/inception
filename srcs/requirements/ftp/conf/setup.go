package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func printColor(color, str string) {
	fmt.Print(color + str + colorReset)
}

func printColorln(color, str string) {
	fmt.Println(color + str + colorReset)
}

func exitIfError(err error) {
	if err != nil {
		printColorln(colorRed, "[FAILED]")
		printColorln(colorRed, err.Error())
		os.Exit(1)
	}
}

func strError(errStr string) {
	printColorln(colorRed, "[FAILED]")
	printColorln(colorRed, errStr)
	os.Exit(1)
}

func validateScriptArgs(args []string) {
	printColor(colorCyan, "[Golang script starting...]")
	if len(args) < 2 {
		strError("invalid number of args, need minimum 2 args")
	}
	printColorln(colorGreen, "[SUCCESS]")
}

func moveFile(from, to string) {
	mv := exec.Command("cp", "-R", from, to)
	err := mv.Run()
	exitIfError(err)
}

func startFtp(ftpBin string) *exec.Cmd {
	printColor(colorCyan, "[start vsftpd...]")
	ftp := exec.Command(ftpBin)
	err := ftp.Start()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	return ftp
}

func createUser(user, password string) {
	printColor(colorCyan, "[create user...]")
	addUser := exec.Command("adduser", "-D", "--home", "/app/data", "--shell", "bash", user)
	err := addUser.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[set password...]")
	setPassword := exec.Command("bash", "-c", fmt.Sprintf("echo %s:%s | chpasswd", user, password))
	err = setPassword.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
}
func main() {
	args := os.Args
	validateScriptArgs(args)
	moveFile("/app/vsftpd.conf", "/etc/vsftpd.conf")
	if _, err := user.Lookup(os.Getenv("FTP_USER")); err != nil {
		createUser(os.Getenv("FTP_USER"), os.Getenv("FTP_PASSWORD"))
	}
	ftp := startFtp(args[1])
	err := ftp.Wait()
	exitIfError(err)
}
