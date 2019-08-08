package main

import (
	"fmt"
	"log"
	"strings"

	Aargh "github.com/jabernardo/aargh"
	Hidden "github.com/jabernardo/msg/hidden"
)

// # Hide
//
// ## Arguments:
//  - app (*Aargh.App) Aargh Application Instance
//
// ## Options:
//  `--src`    `(optional)` Where to hide your message?
//
//  `--out`    `(optional)` Output file
//
//  `--msg`    `(required)` Message to hide
//
// ## Flags:
//  -q       Do not log output
//
// Usage:
//  ### Ouput redirection
//  > $ `msg hide --msg="No pain. No gain." > secret_message.txt`
//
//  ### Output file
//  > $ `msg hide --msg="No pain. No gain." --out="./secret_message.txt"`
//
//  ### From source
//  > $ `msg hide --src="./paragraph.txt" --msg="No pain. No gain." --out="./secret_message.txt"`
//
func command_hide(app *Aargh.App) {
	var src string = app.GetOption("src", "pangram")
	var out string = app.GetOption("out")
	var msg string = app.GetOption("msg")
	var src_content string = "pangram"
	var err error

	if strings.Compare(src, "pangram") != 0 {
		src_content, err = file_get_contents(src)

		if err != nil {
			log.Fatalln("[msg] Invalid source file.")
			return
		}
	}

	content := Hidden.New(src_content, msg)

	if len(out) > 0 {
		err := file_put_contents(out, content)

		if err != nil {
			log.Fatalln("[msg] Invalid source file.")
			return
		}
	}

	if !app.HasFlag("q") {
		fmt.Println(content)
	}
}
