package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"
	"strings"
)

// NewServer creates a new server listening at socketPath
func NewServer(socketPath string) {
	log.Print("Making a new server")
	sock, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatal("Error creating socket:", err)
	}
	defer sock.Close()

	message := make(chan connMsg)

	go handleSocket(message, sock)
	handleMessages(message)
}

type connMsg struct {
	conn    net.Conn
	id      string
	command string
	args    []string
}

func handleMessages(message chan connMsg) {
	for msg := range message {
		switch msg.command {
		// Leave the function, thus quiting the server
		case "quit":
			return

		// Simple ping/pong command
		case "ping":
			msg.conn.Write([]byte("pong\n"))

		default:
			log.Printf("Unknown command from conn(%s): %s", msg.id, msg.command)
		}
	}
}

func handleSocket(message chan connMsg, sock net.Listener) {
	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Fatal("Error accepting connectiong:", err)
		}
		defer conn.Close()

		go handleConnection(message, conn)
	}
}

func handleConnection(message chan connMsg, conn net.Conn) {
	id := newID()

	log.Print("New connection, id: " + id)
	for {
		rd := bufio.NewReader(conn)
		line, err := rd.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				log.Printf("Connection (%s) closed", id)
			} else {
				log.Print("Error while reading from connection:")
				log.Print("  " + err.Error())
			}
			return
		}

		line = strings.TrimSpace(line)
		cmd := strings.Split(line, " ")
		args := cmd[1:]

		message <- connMsg{conn, id, cmd[0], args}
	}
}

func newID() string {
	rb := make([]byte, 4)
	_, err := rand.Read(rb)
	if err != nil {
		log.Fatal("Error generating ID: " + err.Error())
	}

	return hex.EncodeToString(rb)
}
