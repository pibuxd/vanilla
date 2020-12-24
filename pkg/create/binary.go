package create

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"os"
	"os/exec"
)

func Bin() {
	P := t.Package{}
	P.Type = "bin"

	packagePath := ""
	packagePath = "/home/pibu/Vanilla/server/files/main-bin"
	//fmt.Printf("Path to package: ")
	//fmt.Scanf("%s", &packagePath)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Name: ")))
	fmt.Scanf("%s", &P.Name)

	P.Location = "/usr/bin/" + string(P.Name)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)

	file, err := json.Marshal(P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	err = ioutil.WriteFile("/home/pibu/Vanilla/server/files/main-bin.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	cmd := exec.Command("sh", os.Getenv("HOME")+"/vanilla/pkg/scripts/uploadBin.sh", string(packagePath))
	err = cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success: ")), "Uploaded")
}
