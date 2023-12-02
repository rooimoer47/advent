package main

import (
	"fmt"
	"io/ioutil"
  )

func main() {
    advent1a()
}

func adventTemplate() {
	data, err := ioutil.ReadFile("temp.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	content := string(data)
	fmt.Println(content)
}

func advent1a() {
	data, err := ioutil.ReadFile("1b.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//content := string(data)
	// fmt.Println(content)

	digitMap := map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six": 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}

	sum := 0
	first := -1
	last := -1
	for index, r := range data {
        if r >= '0' && r <= '9' {
			if first == -1 {
				first = int(r) - '0'
			}
            last = int(r) - '0'
			//fmt.Println("set nums ", first, last)
        } else if r != '\n'{
			for digit := range digitMap {
				//fmt.Println(digit)
				if index + len(digit) < len(data) && digit == string(data[index:index + len(digit)]) {
					if first == -1 {
						first = digitMap[digit]
					}
					last = digitMap[digit]
				}
			}
		} else {
			sum = sum + first*10 + last
			//fmt.Println("add ", first*10 + last)
			first = -1
		    last = -1
		}
    }
	fmt.Println(sum + first*10 + last)
}