package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"regexp"

	"github.com/codegangsta/cli"
)

var Commands = []cli.Command{
	commandGet,
	commandPut,
	commandDelete,
	commandGetAws,
}

var commandGet = cli.Command{
	Name:  "get",
	Usage: "get <key>",
	Description: `
`,
	Action: doGet,
}

var commandPut = cli.Command{
	Name:  "put",
	Usage: "put <key> <value>",
	Description: `
`,
	Action: doPut,
}

var commandDelete = cli.Command{
	Name:  "delete",
	Usage: "delete <key>",
	Description: `
`,
	Action: doDelete,
}

var commandGetAws = cli.Command{
	Name:  "get-aws",
	Usage: "get-aws desuyo",
	Description: `
`,
	Action: doGetAws,
}

func debug(v ...interface{}) {
	if os.Getenv("DEBUG") != "" {
		log.Println(v...)
	}
}

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func doGet(c *cli.Context) {
	key := c.Args().First()
	if validateKey(key) {
		showUsage(c)
	}

	keypath := getShelfHome() + key
	if _, err := os.Stat(keypath); err == nil {
		data, err := ioutil.ReadFile(keypath)
		if err != nil {
			fmt.Println(string(data))
			os.Exit(1)
		}
		fmt.Print(string(data))
	}
}

func doPut(c *cli.Context) {
	key := c.Args().Get(0)
	value := c.Args().Get(1)
	if validateKey(key) || value == "" {
		showUsage(c)
	}

	keypath := getShelfHome() + key
	dir, _ := path.Split(keypath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// create directory
		os.MkdirAll(dir, 0777)
	}
	// write
	if err := ioutil.WriteFile(keypath, []byte(value), 0600); err != nil {
		fmt.Println(err)
	}
}

func doDelete(c *cli.Context) {
	key := c.Args().First()
	if validateKey(key) {
		showUsage(c)
	}

	keypath := getShelfHome() + key
	if _, err := os.Stat(keypath); err == nil {
		if err := os.Remove(keypath); err != nil {
			fmt.Println(err)
		}
	}
}

func doGetAws(c *cli.Context) {
}

func getShelfHome() string {
	shelf_home := os.Getenv("SHELF_HOME")
	if shelf_home == "" {
		// set userhome directory
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		shelf_home = usr.HomeDir + "/.shelf"
	}
	return shelf_home
}

func validateKey(key string) bool {
	if key == "" || key == "/" || regexp.MustCompile(`^/(\w)+`).MatchString(key) == false || regexp.MustCompile(`(\w)+/$`).MatchString(key) == true {
		return true
	}
	return false
}

func showUsage(c *cli.Context) {
	fmt.Println("Usage:", c.App.Command(c.Command.Name).Usage)
	os.Exit(1)
}
