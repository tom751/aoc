package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	horizontalPos := 0
	depth := 0
	aim := 0
	for scanner.Scan() {
		row := scanner.Text()
		split := strings.Split(row, " ")
		value, err := strconv.Atoi(split[1])

		if err != nil {
			log.Fatal(err)
		}

		switch split[0] {
		case "forward":
			horizontalPos += value
			depth += (aim * value)
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(horizontalPos * depth)
}
