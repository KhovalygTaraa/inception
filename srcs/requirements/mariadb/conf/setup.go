package main

import(
	"os"
	"fmt"
	"strconv"
	"strings"
	"os/exec"
	"os/user"
	"path/filepath"
)

const(
	colorReset = "\033[0m"
	colorRed = "\033[31m"
	colorGreen = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan = "\033[36m"
	colorWhite = "\033[37m"
)

func printColor(color, str string) {
	fmt.Print(color + str + colorReset)
}

func printColorln(color, str string) {
	fmt.Println(color + str + colorReset)
}

func ChownR(path string, uid, gid string) error {
	function := func(name string, info os.FileInfo, err error) error {
		if err != nil  {
			return err
		}
		uidDecimal, err := strconv.Atoi(uid)
		if err != nil  {
			return err
		}
		gidDecimal, err := strconv.Atoi(gid)
		if err != nil  {
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

func prepareMariadbWorkspace(mariadbConf, appDir, dataDir, mysqldDir string){
	printColor(colorCyan, "[Prepare workspace for mariadb...]")
	mysqlUser, err := user.Lookup("mysql")
	exitIfError(err)
	exitIfError(os.MkdirAll(mysqldDir, 0755))
	exitIfError(os.MkdirAll(dataDir, 0755))
	exitIfError(ChownR(appDir, mysqlUser.Uid, mysqlUser.Gid))
	exitIfError(ChownR(mysqldDir, mysqlUser.Uid, mysqlUser.Gid))
	exitIfError(ChownR(dataDir, mysqlUser.Uid, mysqlUser.Gid))
	exitIfError(os.Chmod(mariadbConf, 0644))
	printColorln(colorGreen, "[SUCCESS]")	
}


func start(args []string) {
	printColor(colorCyan, "[Golang script starting...]")
	if len(args) < 2 {
		strError("invalid number of args, need minimun 2 args")
	}
	printColorln(colorGreen, "[SUCCESS]")
}

func checkArgs(args []string) {
	printColor(colorCyan, "[Check args]")
	if len(args) != 4 {
		strError("invalid number of args, need 4 args for mariadb")
	}
	printColorln(colorGreen, "[SUCCESS]")
}

func installMariadb(mariadbConf, dataDir string) {
	printColor(colorCyan, "[Mariadb installing...]")
	mariadbInstallDb := exec.Command("mariadb-install-db", mariadbConf, dataDir)
	err := mariadbInstallDb.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")

}

func startMariadb(args []string) {
	printColor(colorCyan, "[Mariadb starting...]")
	mariadbd := exec.Command(args[1], args[2], args[3])
	err := mariadbd.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
}

func main() {
	args := os.Args
	start(args)
	if args[1] == "mariadbd" {
		printColorln(colorCyan, "[Mariadbd]")
		checkArgs(args)
		mariadbConf := strings.Split(args[2], "=")[1]
		dataDir := strings.Split(args[3], "=")[1]
		prepareMariadbWorkspace(mariadbConf, "/app", dataDir, "/run/mysqld")
		if _, err := os.Stat(dataDir + "/mysql"); os.IsNotExist(err) {
			installMariadb(args[2], args[3])
		}
		startMariadb(args)
	}
}
