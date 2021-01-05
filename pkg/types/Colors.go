package types

import (
	"flag"
	"github.com/logrusorgru/aurora"
)

var c aurora.Aurora
var colors = flag.Bool("colors", false, "enable or disable colors")

func init() {
	flag.Parse()
	c = aurora.NewAurora(*colors)
}
