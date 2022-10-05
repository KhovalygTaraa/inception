package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
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

func moveFile(from, to string) {
	//err := os.Rename(from, to)
	mv := exec.Command("cp", "-R", from, to)
	err := mv.Run()
	exitIfError(err)
}

func moveConfigs() {
	printColor(colorCyan, "[move configs...]")
	moveFile(os.Getenv("PHP_FPM_GLOBAL_CONFIG"), "/etc/php8/php-fpm.conf")
	moveFile(os.Getenv("PHP_FPM_WWW_CONFIG"), "/etc/php8/php-fpm.d/www.conf")
	printColorln(colorGreen, "[SUCCESS]")
}

func givePermissions() {
	printColor(colorCyan, "[give permissions to php-fpm default user...]")
	phpFpmDefaultUser, err := user.Lookup("nobody")
	exitIfError(err)
	exitIfError(ChownR(os.Getenv("APP_DIR"), phpFpmDefaultUser.Uid, phpFpmDefaultUser.Gid))
	printColorln(colorGreen, "[SUCCESS]")

}

func startPhpFpm(phpFpmBin string) *exec.Cmd {
	printColor(colorCyan, "[start php-fpm8...]")
	phpFpm8 := exec.Command(phpFpmBin)
	err := phpFpm8.Start()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	return phpFpm8
}

func getAdminer(link string) {
	var outputFormat string = "adminer.php"

	printColor(colorCyan, "[download adminer...]")
	wget := exec.Command("wget", "-O", outputFormat, link)
	err := wget.Run()
	exitIfError(err)
	moveFile(os.Getenv("APP_DIR")+"/"+outputFormat, os.Getenv("DATA_DIR")+"/"+outputFormat)
	printColorln(colorGreen, "[SUCCESS]")
}

func main() {
	args := os.Args
	validateScriptArgs(args)
	if _, err := os.Stat(os.Getenv("DATA_DIR") + "/adminer.php"); os.IsNotExist(err) {
		getAdminer(os.Getenv("ADMINER_LINK"))
	}
	moveConfigs()
	phpFpm8 := startPhpFpm(args[1])
	givePermissions()
	err := phpFpm8.Wait()
	exitIfError(err)
}
