package Hidden

import (
	"io/ioutil"
	"strings"
)

type Message struct {
	// Source File
	Source 	string
	// Content
	Content string
	// Output File
	Output	string
}

func New(src string, out string, msg string) *Message {
	new_msg := &Message{}
	new_msg.Source = src
	new_msg.Output = out
	new_msg.Content = msg

	return new_msg
}

func (app *Message) GetContents(src string) (string, error) {
	data, err := ioutil.ReadFile(src)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (app *Message) Encrypt() (string, error) {
	var encrypted strings.Builder
	source, err := app.GetContents(app.Source)

	if err != nil {
		return "", err
	}

	last_index := 0;

	for _, char := range source {
		if last_index < len(app.Content) {
			for i := last_index; i < len(app.Content); i++ {

				char_content := strings.ToUpper(string(app.Content[i]))
				char_source := strings.ToUpper(string(char))

				if char_content == char_source {
					encrypted.WriteString(strings.ToUpper(char_source))

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
