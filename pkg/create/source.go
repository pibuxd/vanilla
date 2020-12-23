package create

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"os"
	"os/exec"
)

func Src() {
	P := t.Package{}
	P.Type = "src"

	//packagePath := ""
	//fmt.Printf("Path to package: ")
	//fmt.Scanf("%s", &packagePath)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Name: ")))
	fmt.Scanf("%s", &P.Name)

	P.Location = "/usr/bin/" + string(P.Name)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)
	packagePath := "/home/pibu/Vanilla/server/files/main.tar.gz"

	//command1 := "-X POST http://pibux.pl:2137/upload"
	//command2 := "-F \"file=@" + string(packagePath) + "\""
	//command3 := "-H \"Content-Type: multipart/form-data\""

	//cmd := exec.Command("curl", string(command1), string(command2), string(command3))
	cmd := exec.Command("sh", os.Getenv("HOME")+"/vanilla/pkg/scripts/uploadSrc.sh", string(packagePath))
	err := cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("error: ")), err)
		return
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success: ")), "Uploaded")
}
