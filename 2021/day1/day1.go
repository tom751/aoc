package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var prevNum int
	index := 0
	counter := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if index > 0 {
			if num > prevNum {
				counter++
			}
		}

		index++
		prevNum = num

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(counter)
}
