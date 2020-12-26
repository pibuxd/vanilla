package remove

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"os"
	"os/exec"
)

func DeleteDuplicate(P []t.Package, packageName string) []t.Package {
	T := []t.Package{}

	for _, p := range P {
		if p.Name != packageName {
			T = append(T, t.Package{p.Name, p.Version})
		}
	}
	return T
}

func RemPackage(packageName string) {
	cmd := exec.Command("sh", os.Getenv("HOME")+"/.vanilla/scripts/remove.sh", packageName)
	err := cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	P := []t.Package{}
	data, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.vanilla/data/installed-packages.json")
	json.Unmarshal([]byte(data), &P)

	P = DeleteDuplicate(P, packageName)

	file, err := json.Marshal(P)

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	err = ioutil.WriteFile(os.Getenv("HOME")+"/.vanilla/data/installed-packages.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success:")), c.Bold(c.Blue(packageName)), "has been removed")
}
