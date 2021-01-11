package create

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
	t "github.com/pibuxd/vanilla/pkg/types"
	"os"
	"os/exec"
	"path"
)

// NewPackage creates a new package from tar.gz
func NewPackage() {
	P := t.Package{}

	packagePath := ""
	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Path to package: ")))
	fmt.Scanf("%s", &packagePath)

	P.Name = path.Base(packagePath)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Version: ")))
	fmt.Scanf("%s", &P.Version)

	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Dependiences: ")))
	fmt.Scanf("%s", &P.Depend)

	haslo := ""
	fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("Password: ")))
	fmt.Scanf("%s", haslo)

	f, err := os.Create("/etc/vanilla/tmp/" + P.Name + ".txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror: ")), err)
		return
	}

	_, err = f.WriteString(P.Version)
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror:")), err)
		return
	}

	cmd := exec.Command("sh", "/etc/vanilla/scripts/upload.sh", string(packagePath))
	err = cmd.Run()

	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror:")), err)
		return
	}

	err = os.Remove("/etc/vanilla/tmp/" + P.Name + ".txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("\nerror:")), err)
		return
	}

	fmt.Printf("\n")
	fmt.Println(c.Bold(c.Green("success:")), c.Bold(c.Blue(P.Name)), "has been uploaded")
}
