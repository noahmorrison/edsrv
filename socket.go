package main

import (
	"bufio"
	"net"
	"strings"
)

// Listener listens to the unix socket at
// /tmp/aneditor.sock and runs commands
type Listener struct {
	running bool
}

// ListenTo listens to the socket /tmp/aneditor.sock
// and handles incoming connections
func ListenTo(message chan string, socketPath string) {
	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		println("Error connecting to socket:", err.Error())
		return
	}
	defer sock.Close()

	for {
		conn, err := sock.Accept()
		if err != nil {
			println("Accept error", err.Error())
			return
		}
		defer conn.Close()

		if !handleConnection(message, conn) {
			break
		}
	}
}

// handleConnection takes the message channel, and a connection.
// it returns true if the socket should stay open.
func handleConnection(message chan string, conn net.Conn) bool {
	message <- "alert New connection"
	for {
		rd := bufio.NewReader(conn)
		line, err := rd.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				message <- "alert connection closed"
			} else {
				message <- "alert Error while reading from connection: " +
					err.Error()
			}
			return true
		}

		line = strings.TrimSpace(line)
		if line == "quit" {
			close(message)
			return false
		}
		message <- line
	}
}
