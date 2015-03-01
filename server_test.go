package main

import "testing"

func TestPingPong(t *testing.T) {
	// Connect to the server
	message := make(chan string)
	go ConnectTo(message, "/tmp/aneditor.sock")

	// Ping
	message <- "ping"

	// Pong?
	if <-message != "pong" {
		t.Error("Ping did not pong")
	}
}

func TestQuit(t *testing.T) {
	message := make(chan string)
	go ConnectTo(message, "/tmp/aneditor.sock")

	message <- "quit"

	_, opened := <-message
	if opened {
		t.Error("Server did not quit")
	}
}
