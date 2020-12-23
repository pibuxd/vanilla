package main

import (
	"fmt"
	"github.com/pibuxd/vanilla/pkg/create"
)

func handleExist(packageName string) {
	// check if the package exists
	return
}

func handleCreate() {
	// ask if binary or source package and handle it

	packageType := ""
	fmt.Printf("Type [bin/src]: ")
	fmt.Scanf("%s", &packageType)

	if packageType == "bin" {
		createBin()
	} else if packageType == "src" {
		createSrc()
	} else {
		fmt.Println("[ERROR] Wrong type")
		handleCreate()
	}
}

func handleSync(packageName string) {
	// install package
	return
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
