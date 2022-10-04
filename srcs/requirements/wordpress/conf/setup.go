package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	moveFile(os.Getenv("WP_CONFIG"), os.Getenv("WORDPRESS_PATH")+"/wp-config.php")
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

func getWordpress(link string) {
	var outputFormat string = "wordpress.tar.gz"

	printColor(colorCyan, "[download wordpress...]")
	wget := exec.Command("wget", "-O", outputFormat, link)
	err := wget.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[unzip wordpress...]")
	tar := exec.Command("tar", "-xzf", outputFormat, "-C", os.Getenv("DATA_DIR"))
	err = tar.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
}

func getWpCli(link string) {
	var outputFormat string = "wp-cli.phar"

	printColor(colorCyan, "[download wp-cli...]")
	wget := exec.Command("wget", "-O", outputFormat, link)
	err := wget.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[check is workable...]")
	workable := exec.Command("php8", outputFormat, "--info")
	err = workable.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[chmod wp-cli.phar...]")
	exitIfError(os.Chmod(outputFormat, 0755))
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[move wp-cli.phar to /usr/local/bin/wp...]")
	moveFile(outputFormat, "/usr/local/bin/wp")
	printColorln(colorGreen, "[SUCCESS]")

}

func installWp() {
	url := fmt.Sprintf("--url=https://swquinc.42.fr/wordpress")
	title := fmt.Sprintf("--title=%s", os.Getenv("WP_TITLE"))
	admin := fmt.Sprintf("--admin_user=%s", os.Getenv("WP_ADMIN"))
	adminPass := fmt.Sprintf("--admin_password=%s", os.Getenv("WP_ADMIN_PASSWORD"))
	adminMail := fmt.Sprintf("--admin_email=%s", os.Getenv("WP_ADMIN_MAIL"))
	path := fmt.Sprintf("--path=%s", os.Getenv("WORDPRESS_PATH"))
	printColor(colorCyan, "[install wordpress...]")
	install := exec.Command("wp", path, "core", "install", url, title, admin, adminPass, adminMail)
	err := install.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[set language...]")
	setLang := exec.Command("wp", path, "language", "core", "activate", "en_US")
	err = setLang.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
}

func installRedis() {
	path := fmt.Sprintf("--path=%s", os.Getenv("WORDPRESS_PATH"))
	printColor(colorCyan, "[install redis-cache plugin...]")
	redisInstall := exec.Command("wp", path, "plugin", "install", "redis-cache")
	err := redisInstall.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
	printColor(colorCyan, "[activate redis-cache plugin...]")
	redisActivate := exec.Command("wp", path, "plugin", "activate", "redis-cache")
	err = redisActivate.Run()
	exitIfError(err)
	redisEnable := exec.Command("wp", path, "redis", "enable")
	err = redisEnable.Run()
	exitIfError(err)
	printColorln(colorGreen, "[SUCCESS]")
}

func setUser(user, password, mail string) {
	path := fmt.Sprintf("--path=%s", os.Getenv("WORDPRESS_PATH"))
	userPass := fmt.Sprintf("--user_pass=%s", password)
	role := fmt.Sprintf("--role=editor")

	createUser := exec.Command("wp", path, "user", "create", "redis-cache", user, mail, userPass, role)
	err := createUser.Run()
	exitIfError(err)
}

func main() {
	args := os.Args
	validateScriptArgs(args)
	if _, err := os.Stat(os.Getenv("WORDPRESS_PATH")); os.IsNotExist(err) {
		getWordpress(os.Getenv("WORDPRESS_LINK"))
		moveConfigs()
		getWpCli(os.Getenv("WP_CLI_LINK"))
		installWp()
		setUser(os.Getenv("WP_USER"), os.Getenv("WP_USER_PASSWORD"), os.Getenv("WP_USER_MAIL"))
		installRedis()
	} else {
		moveConfigs()
	}
	phpFpm8 := startPhpFpm(args[1])
	givePermissions()
	err := phpFpm8.Wait()
	exitIfError(err)
}
