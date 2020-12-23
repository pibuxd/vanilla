package main

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	"github.com/pibuxd/vanilla/pkg/create"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"os"
)

func handleExist(packageName string) bool {
	// check if the package exists

	P := []t.Package{}
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/vanilla/data/installed-packages.json")
	if err != nil {
		t.printError()
		fmt.Println(err)
	}

	err = json.Unmarshal([]byte(data), &P)
	if err != nil {
		t.printError()
		fmt.Println(err)
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

	packageType := ""
	fmt.Printf("Type [bin/src]: ")
	fmt.Scanf("%s", &packageType)

	if packageType == "bin" {
		create.Bin()
	} else if packageType == "src" {
		create.Src()
	} else {
		t.printError()
		fmt.Println("Wrong type")
		handleCreate()
	}
}

func handleSync(packageName string) {
	// install package

	if handleExist(packageName) {
		fmt.Printf("Package %s already exists\n", c.Blue(packageName))
		return
	}

	fmt.Printf(c.Sprintf(c.Bold("Installing:: %s\n"), c.Blue(packageName)))
}

func handleRemove(packageName string) {
	// remove package
	return
}

func handleListAll() {
	// print package names with location from json file
	return
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
