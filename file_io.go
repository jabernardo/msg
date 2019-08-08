package main

import (
	"io/ioutil"
)

// file_get_contents - Get file contents
//
// Arguments:
//  - src (string) Source file path
//
// Returns:
//  - (string) File contents
//  - (error)  Error
func file_get_contents(src string) (string, error) {
	data, err := ioutil.ReadFile(src)

	if err != nil {
		return "", err
	}

	return string(data), nil
}

// file_put_contents - Write to file
//
// Arguments:
//  - out (string) Output file
//  - content (string) File content
//
// Returns:
//  - (error) Error
//
func file_put_contents(out string, content string) error {
	msg := []byte(content)
	err := ioutil.WriteFile(out, msg, 0644)
	return err
}
