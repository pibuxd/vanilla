package main

import (
	"os"
)

func main() {
	operation := ""
	if len(os.Args) > 1 {
		operation = os.Args[1]
	}

	packageName := ""
	if len(os.Args) > 2 {
		packageName = os.Args[2]
	}

	switch operation {
	case "h":
		usage()

	case "s":
		handleSync(packageName)

	case "r":
		handleRemove(packageName)

	case "q":
		handleListAll()

	case "e":
		handleExist(packageName)

	case "up":
		handleUpdate()

	case "ug":
		handleUpgrade()

	case "cr":
		handleCreate()

	default:
		usage()
	}

}
