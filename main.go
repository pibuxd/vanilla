// socket client

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1:2137")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Text to send: ")
		text, _ := reader.ReadString('\n')
		// send to server
		fmt.Fprintf(conn, text+"\n")
		// wait for reply
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message form server: " + message)
	}
}
