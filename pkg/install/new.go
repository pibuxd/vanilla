package install

import (
	"encoding/json"
	"fmt"
	c "github.com/logrusorgru/aurora"
	"github.com/pibuxd/vanilla/pkg/remove"
	t "github.com/pibuxd/vanilla/pkg/types"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
)

func MakePackage(packageName string) {
	cmd := exec.Command("sh", os.Getenv("HOME")+"/.vanilla/scripts/download.sh", packageName)
	err := cmd.Run()
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	P := []t.Package{}

	resp, err := http.Get("https://pibux.pl/data/" + packageName + ".tar.gz.txt")
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}
	defer resp.Body.Close()

	versionTemp, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	data, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.vanilla/data/installed-packages.json")
	json.Unmarshal([]byte(data), &P)
	P = remove.DeleteDuplicate(P, packageName)

	P = append(P, t.Package{packageName, string(versionTemp)})

	file, err := json.Marshal(P)

	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}

	err = ioutil.WriteFile(os.Getenv("HOME")+"/.vanilla/data/installed-packages.json", file, 0644)
	if err != nil {
		fmt.Println(c.Bold(c.Red("error:")), err)
	}
}
