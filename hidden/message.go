package Hidden

import (
	"io/ioutil"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Message Structure
type Message struct {
	// Source File
	Source string
	// Content
	Content string
	// Output File
	Output string
}

// Create a new instance of Message
//
// Arguments:
//  - src (string) Source file path
//  - out (string) Output file path
//  - msg (string) Message
//
// Returns:
//  - (*Message) New Message struct
func New(src string, out string, msg string) *Message {
	new_msg := &Message{}
	new_msg.Source = src
	new_msg.Output = out
	new_msg.Content = msg

	return new_msg
}

// (*Message) GetContents - Get file contents
//
// Arguments:
//  - src (string) Source file path
//
// Returns:
//  - (string) File contents
//  - (error)  Error
func (app *Message) GetContents(src string) (string, error) {
	data, err := ioutil.ReadFile(src)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// (*Message) Encrypt - Encrypt message
//
// Returns
//  - (string) Encrypted message
//  - (error)  Error
func (app *Message) Encrypt() (string, error) {
	var encrypted strings.Builder
	source, err := app.GetContents(app.Source)

	if err != nil {
		return "", err
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

	return encrypted.String(), nil
}
