package main

import (
	"fmt"

	Aargh "github.com/jabernardo/aargh"
)

// Help function
//
// Params:
//	app *Aargh.App Aargh Application Instance
//
//
// Returns:
//	void
//
func command_help(app *Aargh.App) {
	fmt.Printf("Hide Message @%s\n\n", app.Version)
	fmt.Printf("Commands:\n")
	fmt.Printf("\thide\tHide message\n")
	fmt.Printf("\t\texample: msg hide --msg=\"Bo pain. No gain.\" > secret_message.txt\n")
	fmt.Printf("\nOptions:\n\n")
	fmt.Printf("\t--msg\tMessage to hide\n")
	fmt.Printf("\t--src\tSource file for encryption\n")
	fmt.Printf("\t--out\tOutput file\n")
	fmt.Printf("\nFlags:\n\n")
	fmt.Printf("\t-q\tDo not log\n")
}
