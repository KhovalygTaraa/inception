package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
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

func prepareMariadbWorkspace(mariadbConf, appDir, dataDir, mysqldDir string) {
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

func validateScriptArgs(args []string) {
	printColor(colorCyan, "[Golang script starting...]")
	if len(args) < 2 {
		strError("invalid number of args, need minimum 2 args")
	}
	printColorln(colorGreen, "[SUCCESS]")
}

func validateMariadbArgs(args []string) {
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

func startMariadb(args []string) *exec.Cmd {
	printColor(colorCyan, "[Mariadb running...]")
	mariadbd := exec.Command(args[1], args[2], args[3])
	err := mariadbd.Start()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	return mariadbd
}

func setRootPassword(rootPassword string) {
	printColor(colorCyan, "[Connect to db...]")
	db, err := sql.Open("mysql", "root@unix(/run/mysqld/mysqld.sock)/")
	defer db.Close()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	for db.Ping() != nil {
	}
	printColorln(colorCyan, "[Set root password...]")
	_, err = db.Exec(fmt.Sprintf("ALTER USER 'root'@'localhost' IDENTIFIED BY \"%s\"", rootPassword))
	_, err = db.Exec(fmt.Sprintf("FLUSH PRIVILEGES"))
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	//rows, err := db.Query("SELECT user, host, password FROM mysql.user WHERE user = 'root'")
	//defer rows.Close()
	//exitIfError(err)
	//for rows.Next() {
	//	var user, host, password string
	//	exitIfError(rows.Scan(&user, &host, &password))
	//	printColorln(colorPurple, user+" | "+host+" | "+password)
	//}
}

func main() {
	args := os.Args
	isFirstStart := false

	validateScriptArgs(args)
	if args[1] == "mariadbd" {
		printColorln(colorCyan, "[Start mariadbd]")
		validateMariadbArgs(args)
		mariadbConf := strings.Split(args[2], "=")[1]
		dataDir := strings.Split(args[3], "=")[1]
		prepareMariadbWorkspace(mariadbConf, "/app", dataDir, "/run/mysqld")
		if _, err := os.Stat(dataDir + "/mysql"); os.IsNotExist(err) {
			installMariadb(args[2], args[3])
			isFirstStart = true
		}
		mariadb := startMariadb(args)
		if isFirstStart {
			setRootPassword(os.Getenv("MARIADB_ROOT_PASSWORD"))
			//prepareDbForWp(os.Getenv("MARIADB_WP_DB"), os.Getenv("MARIADB_WP_USER"), os.Getenv("MARIADB_WP_PASSWORD"))
		}
		//printColorln(colorBlue, "Mariadb pid = ")
		//printColor(colorGreen, strconv.Itoa(mariadb.Process.Pid))
		err := mariadb.Wait()
		exitIfError(err)
	}
}
