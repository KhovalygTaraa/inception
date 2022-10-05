package main

import (
	"fmt"
	"os"
	"os/exec"
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

func startPrometheus(prometheusBin, config string) *exec.Cmd {
	printColor(colorCyan, "[start prometheus...]")
	ftp := exec.Command(prometheusBin, config)
	err := ftp.Start()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	return ftp
}

func main() {
	args := os.Args
	validateScriptArgs(args)
	ftp := startPrometheus(args[1], args[2])
	err := ftp.Wait()
	exitIfError(err)
}
