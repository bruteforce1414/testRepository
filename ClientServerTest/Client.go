package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {

	conn, _ := net.Dial("tcp", "127.0.0.1:4547")
	for {
				message, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Print("Message from server: "+message)


	}

 fmt.Scanln()
}


