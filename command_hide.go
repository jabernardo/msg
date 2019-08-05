package main

import (
	"fmt"
	"log"

	Aargh "github.com/jabernardo/aargh"
	Message "github.com/jabernardo/msg/hidden"
)

func command_hide(app *Aargh.App) {
	src := app.GetOption("src")
	out := app.GetOption("out")
	msg := app.GetOption("msg")
	hid_msg := Message.New(src, out, msg)

	content, err := hid_msg.Encrypt()

	if err != nil {
		log.Fatalln("[msg] Invalid source file.")
		return
	}

	fmt.Println(content)
}
