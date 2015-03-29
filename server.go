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

	inChan := make(chan connMsg)

	go handleSocket(inChan, sock)
	handleMessages(inChan)
}

type connMsg struct {
	conn net.Conn
	id   string
	text string
}

func handleMessages(message chan connMsg) {
	for msg := range message {
		line := msg.text
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
			log.Printf("Unknown command from conn(%s): %s", msg.id, msg.text)
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

func handleSocket(inChan chan connMsg, sock net.Listener) {
	for {
		conn, err := sock.Accept()
		if err != nil {
			log.Fatal("Error accepting connectiong:", err)
		}
		defer conn.Close()

		go handleConnection(inChan, conn)
	}
}

func handleConnection(inChan chan connMsg, conn net.Conn) {
	id := newID()
	rd := bufio.NewReader(conn)

	log.Print("New connection, id: " + id)
	for {
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
		if cmd != "" {
			inChan <- connMsg{conn, id, cmd}
		}
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
