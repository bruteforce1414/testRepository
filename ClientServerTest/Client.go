package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	    go func() {
    	for {

		conn, err:= net.Dial("tcp", "127.0.0.1:4547")
		    if err != nil {
			fmt.Println(err)
			break;
			//return
			}
				message, _ := bufio.NewReader(conn).ReadString('\n')
				fmt.Print("Message from server: "+message)
				time.Sleep(1 * time.Second)}
	}()

	sigsClient := make(chan os.Signal, 1)
	doneClient := make(chan bool, 1)
	signal.Notify(sigsClient, syscall.SIGINT, syscall.SIGTERM)
	handleClient(sigsClient, doneClient)

}

func handleClient(sigs chan os.Signal,done chan bool) {
	s := <-sigs // ожидание сигнала
	fmt.Println("Программа завершится через 3 секунды. Возникшее событие:", s)
	time.Sleep(3 * time.Second)
	done <- true
	os.Exit(0)

}

