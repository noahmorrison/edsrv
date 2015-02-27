package main

func main() {
	message := make(chan string)
	go ListenTo(message, "/tmp/aneditor.sock")

	RunUI(message)
}
