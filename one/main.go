package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	txtF, err := ioutil.ReadFile("elfs.txt")
	if err != nil {
		log.Fatal(err)
	}
	m := 0
	text := string(txtF) // Convert []byte to string
	elves := strings.Split(text, "\n\n")
	for _, e := range elves {
		elfCals := 0
		for _, cal := range strings.Split(e, "\n") {
			elfCal, err := strconv.Atoi(cal)
			if err != nil {
				fmt.Print(err)
			}
			elfCals += elfCal
		}
		if elfCals > m {
			m = elfCals
		}
	}
	fmt.Printf("Big elf: %d", m)
}
