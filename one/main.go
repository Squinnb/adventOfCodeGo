package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	txtF, err := ioutil.ReadFile("elfs.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(txtF) // Convert []byte to string
	elves := strings.Split(text, "\n\n")
	one(elves)
	two(elves)
}
func one(elfslc []string) {
	m := 0
	for _, e := range elfslc {
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
	fmt.Printf("Big elf: %d \n", m)
}
func two(elfslc []string) {
	m := []int{0, 0, 0}
	for _, e := range elfslc {
		elfCals := 0
		for _, cal := range strings.Split(e, "\n") {
			elfCal, err := strconv.Atoi(cal)
			if err != nil {
				fmt.Print(err)
			}
			elfCals += elfCal
		}
		if elfCals > m[2] {
			m[2] = elfCals
			sort.Ints(m)
		} else if elfCals > m[1] {
			m[1] = elfCals
			sort.Ints(m)
		} else if elfCals > m[0] {
			m[0] = elfCals
			sort.Ints(m)
		}
	}
	sum := m[0] + m[1] + m[2]
	fmt.Printf("Big elf two: %d", sum)
}
