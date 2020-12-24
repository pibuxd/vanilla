package main

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	"github.com/pibuxd/vanilla/pkg/create"
	//"github.com/pibuxd/vanilla/pkg/install"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"os"
	"strings"
)

func handleExist(packageName string) bool {
	// check if the package exists

	P := []t.Package{}
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error: ")), err)
	}

	err = json.Unmarshal([]byte(data), &P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error: ")), err)
	}

	for _, p := range P {
		if p.Name == packageName {
			return true
		}
	}

	return false
}

func handleCreate() {
	// ask if binary or source package and handle it

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Type [bin/src]: ")))
	packageType := ""
	fmt.Scanf("%s", &packageType)

	if strings.Contains(packageType, "b") {
		create.Bin()
	} else if strings.Contains(packageType, "s") {
		create.Src()
	} else {
		fmt.Println(c.Bold(c.Red("error:")), "Wrong type")
		handleCreate()
	}
}

func handleSync(packageName string) {
	// install package

	if handleExist(packageName) {
		fmt.Printf(c.Sprintf(c.Bold(c.Red("error: ")))+"Package %s already exists\n", c.Bold(c.Magenta(packageName)))
		return
	}
	fmt.Printf("resolving dependencies...\nlooking for conflicting packages...\n\n")
	fmt.Printf(c.Sprintf(c.Bold("Package %s\n\n"), c.Bold(c.Magenta(packageName))))
	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Proceed with installation? [Y/n]: ")))
	ifInst := ""
	fmt.Scanf("%s", &ifInst)

	if strings.Contains(ifInst, "n") {
		return
	}
	//install.installPackage(packageName)
}

func handleRemove(packageName string) {
	// remove package
	return
}

func handleListAll() {
	// print package names with location from json file
	P := []t.Package{}

	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error: ")), err)
	}

	err = json.Unmarshal([]byte(data), &P)

	for _, p := range P {
		fmt.Println(c.Bold(c.Blue(p.Name)), c.Bold(p.Version))
	}
}

func handleUpdate() {
	// update all installed packages
	return
}

func handleUpgrade() {
	//upgrade all installed packages
	return
}

func usage() {
	fmt.Println(`Usage:
  vanilla
  vanilla <operation> [...]
  
operations:
  vanilla h              ==> help
  vanilla s <package(s)> ==> sync package(s)
  vanilla r <package(s)  ==> sync package(s)
  vanilla up             ==> update all packages
  vanilla ug             ==> upgrade all packages
  vanilla se <package(s) ==> search for package(s)
  vanilla q              ==> list all installed packages
  vanilla e <package(s)  ==> search for already installed package(s)`)
}
