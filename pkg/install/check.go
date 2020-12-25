package install

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
	"net/http"
	"strings"
)

func firstCheck(packageName string) {
	packageBin := "https://pibux.pl/data/" + packageName + ".json"
	packageSrc := "https://pibux.pl/data/" + packageName + ".tar.gz.json"

	bin := true
	src := true

	resp, err := http.Get(packageBin)
	if err != nil {
		bin = false
	}

	resp, err = http.Get(packageSrc)
	if err != nil {
		src = false
	}

	if !bin && !src {
		fmt.Println(c.Bold(c.Red("error:")), "package", c.Bold(c.Blue(packageName)), "doesn't exist")

	} else if bin && !src {
		installBin(packageName)

	} else if !bin && src {
		installSrc(packageName)

	} else {
		fmt.Printf(c.Sprintf(c.Bold(c.Magenta(":: "))) + c.Sprintf(c.Bold("There are 2 providers available for ")) + c.Sprintf(c.Bold(c.Blue(packageName))) + c.Sprintf(c.Bold("[BIN/src]: ")))

		ifInst := ""
		fmt.Scanf("%s", &ifInst)

		if strings.Contains(ifInst, "b") {
			installBin(packageName)
		} else {
			installSrc(packageName)
		}
	}
}
