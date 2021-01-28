package main

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	"github.com/pibuxd/vanilla/pkg/create"
	"github.com/pibuxd/vanilla/pkg/install"
	"github.com/pibuxd/vanilla/pkg/remove"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)


// Exists check if the package exists
func Exists(packageName string) bool {

	P := []t.Package{}
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "vanilla is broken, can't do anything about it")
	}

	err = json.Unmarshal([]byte(data), &P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	for _, p := range P {
		if p.Name == packageName {
			return true
		}
	}

	return false
}

// upload package to server
func handleCreate() {
	create.NewPackage()
}

// install package
func handleSync(packageName string) {
	packageURL := "https://pibux.pl/data/" + packageName + ".tar.gz.txt"

	resp, err := http.Get(packageURL)
	if resp.StatusCode != http.StatusOK {
		fmt.Println(c.Bold(c.Red("error:")), "package", c.Bold(c.Blue(packageName)), "doesn't exist")
		return

	}
	if Exists(packageName) {
		fmt.Println("Package", c.Bold(c.Blue(packageName)), c.Bold(c.Yellow("warning:")), c.Yellow("reinstalling"), "\n")
	} else {
		fmt.Println("Package", c.Bold(c.Blue(packageName)), "\n")
	}

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Proceed with installation? ")) + c.Sprintf(c.Bold("[Y/n]: ")))

	ifInst := ""
	fmt.Scanf("%s", &ifInst)

	if strings.Contains(ifInst, "y") || ifInst == "" || strings.Contains(ifInst, "Y") {
		install.MakePackage(packageName)
	}
}

// remove package
func handleRemove(packageName string) {
	if !Exists(packageName) {
		fmt.Println(c.Bold(c.Red("error:")), "package", c.Bold(c.Blue(packageName)), "was not found")
		return
	}

	fmt.Println("Package", c.Bold(c.Blue(packageName)), "\n")

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Do you want to remove these packages? ")) + c.Sprintf(c.Bold("[Y/n]: ")))

	ifInst := ""
	fmt.Scanf("%s", &ifInst)

	if strings.Contains(ifInst, "y") || ifInst == "" || strings.Contains(ifInst, "Y") {
		remove.RemPackage(packageName)
	}

}

// print package names with location from json file
func handleListAll() {
	P := []t.Package{}

	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "vanilla is broken, can't do anything about it")
	}

	err = json.Unmarshal([]byte(data), &P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	fmt.Println(c.Bold(c.Green(len(P))), "installed packages\n")

	for _, p := range P {
		fmt.Println(c.Bold(c.Blue(p.Name)), c.Bold(p.Version))
	}
}

func handleExist(packageName string) {
	P := []t.Package{}
	data, err := ioutil.ReadFile(os.Getenv("HOME") + "/.vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "vanilla is broken, can't do anything about it")
	}

	err = json.Unmarshal([]byte(data), &P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	for _, p := range P {
		if p.Name == packageName {
			fmt.Println(c.Bold(c.Blue(p.Name)), c.Bold(p.Version))
			return
		}
	}

	fmt.Println(c.Bold(c.Red("error:")), "package", c.Bold(c.Blue(packageName)), "was not found")
}

// update all installed packages
func handleUpdate() {
	return
}

//upgrade all installed packages
func handleUpgrade() {
	return
}

func usage() {
	name := os.Args[0]
	name = path.Base(name)

	fmt.Printf(`Usage:
  %s
  %s <operation> [...]
  
operations:
  %s h               ==> Display Usage
  %s s  <package(s)> ==> Sync package(s)
  %s r  <package(s)> ==> Remove package(s)
  %s up              ==> Update all packages
  %s ug              ==> Upgrade all packages
  %s se <package(s)> ==> Search for package(s)
  %s q               ==> List all installed packages
  %s e  <package(s)> ==> Search in installed package(s)
  `, name, name, name, name, name, name, name, name, name, name)
}
