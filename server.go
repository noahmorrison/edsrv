package main

import (
	"bufio"
	"crypto/rand"
	"encoding/hex"
	"log"
	"net"
	"strconv"
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
		line := msg.command
		a1, line := getAddress(line)
		line = strings.TrimPrefix(line, ",")
		a2, cmd := getAddress(line)
		log.Printf("a1: %d", a1)
		log.Printf("a2: %d", a2)
		log.Printf("cmd: %s", cmd)

		if cmd == "q" {
			return
		} else if cmd == "Q" {
			return
		} else if strings.HasPrefix(cmd, "e") {
			filename := strings.TrimPrefix(cmd, "e ")
			log.Printf("Editing file: " + filename)
		} else {
			log.Printf("Unknown command from conn(%s): %s", msg.id, msg.command)
		}
	}
}

// getAddress parses a digit from the from of a line.
// It returns the digit, and the rest of the line.
// It returns -1 if no digit was found.
func getAddress(line string) (int, string) {
	offset := len(line)
	for {
		digit, err := strconv.Atoi(line[0:offset])
		if err == nil {
			return digit, line[offset:]
		}
		offset--
		if offset == 0 {
			return -1, line
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
