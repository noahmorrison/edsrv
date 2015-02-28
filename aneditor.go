package main

import (
	"flag"
)

func main() {
	var server bool
	flag.BoolVar(&server, "server", false, "If aneditor should be anserver")
	flag.Parse()

	if server {
		NewServer("/tmp/aneditor.sock")
	} else {
		message := make(chan string)
		go ConnectTo(message, "/tmp/aneditor.sock")

		// RunUI(message)
	}
}
