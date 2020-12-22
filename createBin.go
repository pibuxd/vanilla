package main

import (
	"fmt"
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
	//packagePath := "/home/pibu/Vanilla/server/files/main-bin"

	//command := "curl -X POST http://pibux.pl:2137/upload -F \"upload[]=@" + string(packagePath) + "\" -H \"Content-Type: multipart/form-data\""
  c := exec.Command("ls /home")
  err := c.Run()
  
  if err != nil {
    fmt.Println(err)
  }
}
