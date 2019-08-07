package main

import (
	"fmt"
	"log"

	Aargh "github.com/jabernardo/aargh"
	Hidden "github.com/jabernardo/msg/hidden"
)

func command_hide(app *Aargh.App) {
	var src string = app.GetOption("src", "pangram")
	var out string = app.GetOption("out")
	var msg string = app.GetOption("msg")

	content, err := Hidden.New(src, out, msg)

	if err != nil {
		log.Fatalln("[msg] Invalid source file.")
		return
	}

	fmt.Println(content)
}
