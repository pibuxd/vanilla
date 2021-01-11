package install

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	"github.com/pibuxd/vanilla/pkg/remove"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"net/http"
	"os/exec"
)

// MakePackage creates package from tar.gz on computer
func MakePackage(packageName string) {
	cmd := exec.Command("sh", "/etc/vanilla/scripts/download.sh", packageName)
	err := cmd.Run()
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "could't connect to the server")
	}

	P := []t.Package{}

	resp, err := http.Get("https://pibux.pl/data/version/" + packageName + ".tar.gz.txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "couldn't connect to the server")
	}
	defer resp.Body.Close()

	versionTemp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	resp, err = http.Get("https://pibux.pl/data/depend/" + packageName + ".tar.gz.txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "couldn't connect to the server")
	}
	defer resp.Body.Close()

	versionTemp, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	data, err := ioutil.ReadFile("/etc/vanilla/data/installed-packages.json")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), "vanilla is broken, can't do anything about it")
	}
	
	json.Unmarshal([]byte(data), &P)
	P = remove.DeleteDuplicate(P, packageName)

	P = append(P, t.Package{packageName, string(versionTemp)})

	file, err := json.Marshal(P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	err = ioutil.WriteFile("/etc/vanilla/data/installed-packages.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success:")), c.Bold(c.Blue(packageName)), "has been installed")
}
