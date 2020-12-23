package types

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
)

func printError() {
	fmt.Printf(c.Sprintf(c.Bold(c.Red("error: "))))
}
