package create

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strings"
)

func Src() {
	P := t.Package{}
	P.Type = "src"

	packagePath := ""
	packagePath = "/home/pibu/Vanilla/server/files/main.tar.gz"
	//fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Path to package: ")))
	//fmt.Scanf("%s", &packagePath)

	P.Name = strings.TrimSuffix(strings.TrimSuffix(path.Base(packagePath), path.Ext(path.Base(packagePath))), path.Ext(path.Base(strings.TrimSuffix(path.Base(packagePath), path.Ext(path.Base(packagePath))))))

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Location of compiled binary: ")))
	fmt.Scanf("%s", &P.Location)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)

	cmd := exec.Command("sh", os.Getenv("HOME")+"/vanilla/pkg/scripts/upload.sh", string(packagePath))
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
