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

	var rows []string
	for scanner.Scan() {
		row := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		rows = append(rows, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	length := len(rows[0])

	var gamma string
	var epsilon string
	for i := 0; i < length; i++ {
		zeroCount := 0
		oneCount := 0

		for _, row := range rows {
			char := row[i]
			if char == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}

		if zeroCount > oneCount {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	gammaValue, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonValue, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println(gammaValue * epsilonValue)
}
