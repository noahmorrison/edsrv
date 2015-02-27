package main

import (
	gc "github.com/gbin/goncurses"
	"strings"
	"time"
)

// RunUI creates a new user interface
func RunUI(message chan string) {
	ui := newUI()
	go ui.mainloop()
	defer ui.end()

	for msg := range message {
		cmd := strings.Split(msg, " ")[0]
		args := strings.Split(msg, " ")[1:]

		switch cmd {
		case "alert":
			ui.alert(strings.Join(args, " "))
		default:
			ui.alert("Unknown command: " + cmd)
		}
	}
}

type ui struct {
	status string
	scr    *gc.Window
}

func newUI() ui {
	scr, err := gc.Init()
	if err != nil {
		panic("Could not initialize curses")
	}

	gc.Echo(false)
	gc.Cursor(0)

	ui := ui{"andeditor", scr}
	ui.update()
	return ui
}

// mainloop handles keypress events from the user
// right now it's just a busy/sleep loop
func (ui *ui) mainloop() {
	for {
		time.Sleep(time.Second)
	}
}

// alert displays a message at the bottom
// of the screen for five seconds
func (ui *ui) alert(msg string) {
	row, _ := ui.scr.MaxYX()

	// Print the message at the
	// bottom of the screen
	ui.scr.Move(row-1, 0)
	ui.scr.ClearToEOL()
	ui.scr.Print(msg)
	ui.update()

	// After five seconds,
	// remove the message
	go func() {
		time.Sleep(time.Second * 5)
		ui.scr.Move(row-1, 0)
		ui.scr.ClearToEOL()
		ui.update()
	}()
}

// update redraws important parts of the screen
// and calls refresh on the stdscr
func (ui *ui) update() {
	row, col := ui.scr.MaxYX()
	x := col - len(ui.status)

	ui.scr.Move(row-1, x)
	ui.scr.Print(ui.status)
	ui.scr.Refresh()
}

// end deinitializes aneditor
func (ui *ui) end() {
	ui.alert("Quiting...")
	time.Sleep(time.Second / 2)

	gc.End()
}
