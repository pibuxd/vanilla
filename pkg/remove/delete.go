package remove

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/exec"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
)

// DeleteDuplicate deletes duplicate of a package on computer
func DeleteDuplicate(P []t.Package, packageName string) []t.Package {
	T := []t.Package{}

	for _, p := range P {
		if p.Name != packageName {
			T = append(T, t.Package{p.Name, p.Version, p.Depend})
		}
	}
	return T
}

// RemPackage removes package form computer
func RemPackage(packageName string) {
	cmd := exec.Command("sh", "/etc/vanilla/scripts/remove.sh", packageName)
	err := cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "couldn't connect to the server")
	}

	P := []t.Package{}

	data, _ := ioutil.ReadFile("/etc/vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "vanilla is broken, can't do anything about it")
	}

	json.Unmarshal([]byte(data), &P)

	P = DeleteDuplicate(P, packageName)

	file, err := json.Marshal(P)

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	err = ioutil.WriteFile("/etc/vanilla/data/installed-packages.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success:")), c.Bold(c.Blue(packageName)), "has been removed")
}
