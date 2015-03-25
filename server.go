package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"
	"strings"
)

// NewServer creates a new server listening on port
func NewServer(port string) {
	log.Print("Making a new server")
	sock, err := net.Listen("tcp", ":"+port)
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
}

func handleMessages(message chan connMsg) {
	for msg := range message {
		if msg.command == "q" {
			return
		} else if msg.command == "Q" {
			return
		} else {
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

		cmd := strings.TrimSpace(line)
		message <- connMsg{conn, id, cmd}
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
