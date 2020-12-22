package main

import (
	"encoding/json"
	"fmt"
)

type Package struct {
	Name     string
	Type     string
	Version  string
	Location string
}

func handleExist(packageName string) {
	// check if the package exists
	return
}

func handleCreate() {

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
