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

func Src() {
	P := t.Package{}
	P.Type = "src"

	packagePath := ""
	packagePath = "/home/pibu/Vanilla/server/files/main.tar.gz"
	//fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Path to package: ")))
	//fmt.Scanf("%s", &packagePath)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Name: ")))
	fmt.Scanf("%s", &P.Name)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Location of compiled binary: ")))
	fmt.Scanf("%s", &P.Location)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)

	cmd := exec.Command("sh", os.Getenv("HOME")+"/vanilla/pkg/scripts/uploadSrc.sh", string(packagePath))
	err := cmd.Run()

	file, err := json.Marshal(P)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	err = ioutil.WriteFile("/home/pibu/Vanilla/server/files/main.tar.gz.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success: ")), "Uploaded")
}
