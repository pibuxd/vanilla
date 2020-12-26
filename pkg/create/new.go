package create

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"os"
	"os/exec"
	"path"
)

func NewPackage() {
	P := t.Package{}

	packagePath := ""
	packagePath = "/home/pibu/Vanilla/server/files/main.tar.gz"
	//fmt.Printf("Path to package: ")
	//fmt.Scanf("%s", &packagePath)

	P.Name = path.Base(packagePath)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)

	f, err := os.Create(packagePath + ".txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	_, err = f.WriteString(P.Version)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	cmd := exec.Command("sh", os.Getenv("HOME")+"/.vanilla/scripts/upload.sh", string(packagePath))
	err = cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success: ")), "Uploaded")
}
