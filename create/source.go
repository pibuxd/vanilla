package create

import (
	"fmt"
	"os/exec"
)

func createSrc() {
	P := Package{}
	P.Type = "src"

	//packagePath := ""
	//fmt.Printf("Path to package: ")
	//fmt.Scanf("%s", &packagePath)

	fmt.Printf("Name: ")
	fmt.Scanf("%s", &P.Name)

	P.Location = "/usr/bin/" + string(P.Name)

	fmt.Printf("Version: ")
	fmt.Scanf("%s", &P.Version)
	packagePath := "/home/pibu/Vanilla/server/files/main.tar.gz"

	//command1 := "-X POST http://pibux.pl:2137/upload"
	//command2 := "-F \"file=@" + string(packagePath) + "\""
	//command3 := "-H \"Content-Type: multipart/form-data\""

	//c := exec.Command("curl", string(command1), string(command2), string(command3))
	c := exec.Command("sh", "scripts/uploadSrc.sh", string(packagePath))
	err := c.Run()

	if err != nil {
		fmt.Println(err)
	}
}
