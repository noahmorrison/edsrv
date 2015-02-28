package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

// ConnectTo connects to the socket /tmp/aneditor.sock
func ConnectTo(message chan string, socketPath string) {
	info, err := os.Stat(socketPath)
	if err != nil {
		log.Fatal("Error stating socket:", err.Error())
	}

	if info.Mode().String()[0] != 'S' {
		println(info.Mode().String())
		log.Fatal(socketPath + " exists, but is not a socket!")
	}

	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Fatal("Error connecting to socket: " + err.Error())
	}

	bufconn := bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	bufconn.WriteString("ping\n")
	bufconn.Flush()
	line, err := bufconn.ReadString('\n')
	if err != nil {
		log.Fatal("Error reading line from socket: " + err.Error())
	}
	message <- "alert " + line
}
