package Hidden

import (
	"math/rand"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
)

// Pangrams Collection
var PANGRAMS = [8]string{
	"Jim quickly realized that the beautiful gowns are expensive.",
	"Quirky spud boys can jam after zapping five worthy Polysixes.",
	"Jackie will budget for the most expensive zoology equipment.",
	"Zack Gappow saved the job requirement list for the six boys.",
	"Zelda might fix the job growth plans very quickly on Monday.",
	"The wizard quickly jinxed the gnomes before they vaporized.",
	"Bobby Klun awarded Jayme sixth place for her very high quiz.",
	"All questions asked by five watched experts amaze the judge.",
}

// Message Structure
type Message struct {
	// Source File
	Source string
	// Content
	Content string
}

// Create a new instance of Message
//
// Arguments:
//  - src (string) Source file path (set value to `pangram` for random pangram source)
//  - msg (string) Message
//
// Returns:
//  - (string) Encrypted message
func New(src string, msg string) string {
	new_msg := &Message{}
	new_msg.Source = src
	new_msg.Content = msg

	return new_msg.encrypt()
}

// (*Message).getRandomPangram - Generate default source
//
// Returns:
//  - (string) Default source message
func (app *Message) getRandomPangram() string {
	rand_src := rand.NewSource(time.Now().UnixNano())
	rand := rand.New(rand_src)
	pangram_index := rand.Intn(8)

	default_msg := PANGRAMS[pangram_index]
	var build_msg strings.Builder
	max := len(app.Content)

	for i := 0; i < max; i++ {
		build_msg.WriteString(default_msg)
		build_msg.WriteString(" ")
	}

	return build_msg.String()
}

// (*Message).encrypt - Encrypt message
//
// Returns
//  - (string) Encrypted message
func (app *Message) encrypt() string {
	var encrypted strings.Builder

	source := app.Source

	if strings.Compare(app.Source, "pangram") == 0 {
		source = app.getRandomPangram()
	}

	last_index := 0

	for _, char := range source {
		if last_index < len(app.Content) {
			for i := last_index; i < len(app.Content); i++ {

				char_content := strings.ToUpper(string(app.Content[i]))
				char_source := strings.ToUpper(string(char))
				char_unicode, _ := utf8.DecodeRuneInString(char_content)

				if strings.Compare(char_content, char_source) == 0 || !unicode.IsLetter(char_unicode) {
					encrypted.WriteString(strings.ToUpper(char_content))

					last_index += 1
					break
				} else {
					encrypted.WriteString(strings.ToLower(char_source))
					break
				}
			}
		} else {
			encrypted.WriteString(string(char))
		}
	}

	return encrypted.String()
}
