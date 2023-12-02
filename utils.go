package main

import (
	"fmt"
	"os"
  )

  func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	content := string(data)
	return content, nil
}