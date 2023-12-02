package main

import (
	"fmt"
	"strings"
	"strconv"
)

func advent2a() {
	limitMap := map[string]int{
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	sum := 0
	content, _ := readFile("2a.txt")
	fmt.Println(content)
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		gameid, err := strconv.Atoi(strings.Split(strings.Split(line, ":")[0], " ")[1])
		if err != nil {
			fmt.Println("Error:", err)
		  }
		//fmt.Println(gameid)
		gamebool := true
		sets := strings.Split(strings.TrimSpace(strings.Split(line, ":")[1]), ";")
		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				digit, err := strconv.Atoi(strings.Split(strings.TrimSpace(cube), " ")[0])
				if err != nil {
					fmt.Println("Error:", err)
				  }
				colour := strings.Split(strings.TrimSpace(cube), " ")[1]
				if digit > limitMap[colour] {
					gamebool = false
				}
			}
			if !gamebool {
				break
			}
		}
		if gamebool {
			sum += gameid
		}
	}
	fmt.Println(sum)
}
