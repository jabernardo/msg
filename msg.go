package main

import (
	"log"

	Aargh "github.com/jabernardo/aargh"
)

func main() {
	// Reset log settings
	log.SetFlags(0)

	app := Aargh.New()
	
	app.Command("default", command_help)
	app.Command("hide", command_hide)

	app.Run()
}

