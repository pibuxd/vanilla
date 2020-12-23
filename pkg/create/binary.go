package create

import (
	"fmt"
	"github.com/pibuxd/vanilla/pkg/type"
	"os/exec"
)

func createBin() {
	P := Package{}
	P.Type = "bin"

	//packagePath := ""
	//fmt.Printf("Path to package: ")
	//fmt.Scanf("%s", &packagePath)

	fmt.Printf("Name: ")
	fmt.Scanf("%s", &P.Name)

	P.Location = "/usr/bin/" + string(P.Name)

	fmt.Printf("Version: ")
	fmt.Scanf("%s", &P.Version)
	packagePath := "/home/pibu/Vanilla/server/files/main-bin"

	//command1 := "-X POST http://pibux.pl:2137/upload"
	//command2 := "-F \"file=@" + string(packagePath) + "\""
	//command3 := "-H \"Content-Type: multipart/form-data\""

	//c := exec.Command("curl", string(command1), string(command2), string(command3))
	c := exec.Command("sh", "scripts/uploadBin.sh", string(packagePath))
	err := c.Run()

	if err != nil {
		fmt.Println(err)
	}
}
