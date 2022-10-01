package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
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

func ChownR(path string, uid, gid string) error {
	function := func(name string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		uidDecimal, err := strconv.Atoi(uid)
		if err != nil {
			return err
		}
		gidDecimal, err := strconv.Atoi(gid)
		if err != nil {
			return err
		}
		err = os.Chown(name, uidDecimal, gidDecimal)
		return err
	}
	err := filepath.Walk(path, function)
	return err
}

func strError(errStr string) {
	printColorln(colorRed, "[FAILED]")
	printColorln(colorRed, errStr)
	os.Exit(1)
}

func exitIfError(err error) {
	if err != nil {
		printColorln(colorRed, "[FAILED]")
		printColorln(colorRed, err.Error())
		os.Exit(1)
	}
}

func validateScriptArgs(args []string) {
	printColor(colorCyan, "[Golang script starting...]")
	if len(args) < 2 {
		strError("invalid number of args, need minimum 2 args")
	}
	printColorln(colorGreen, "[SUCCESS]")
}

func copyConfigs() {

}

func startPhpFpm(phpFpmBin string) {
	printColorln(colorCyan, "[start php-fpm8...]")
	phpFpm8 := exec.Command(phpFpmBin)
	err := phpFpm8.Start()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	err = phpFpm8.Wait()
	exitIfError(err)
}

func main() {
	args := os.Args
	validateScriptArgs(args)
	copyConfigs()
	startPhpFpm(args[1])
}
