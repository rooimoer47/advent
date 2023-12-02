package main

import (
	"fmt"
	"io/ioutil"
  )

  func readFile(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	content := string(data)
	return content, nil
}