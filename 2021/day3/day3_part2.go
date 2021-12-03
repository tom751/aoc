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

	oxygenRating := findNum(rows, 0, true)
	scrubberRating := findNum(rows, 0, false)

	oxygenValue, _ := strconv.ParseInt(oxygenRating, 2, 64)
	scrubberValue, _ := strconv.ParseInt(scrubberRating, 2, 64)
	fmt.Println(oxygenValue * scrubberValue)
}

func findNum(rows []string, charIndex int, useHigher bool) string {
	if len(rows) == 1 {
		return rows[0]
	}

	length := len(rows[0])
	if length == charIndex {
		return ""
	}

	var zeroes []string
	var ones []string

	for _, row := range rows {
		char := row[charIndex]
		if char == '0' {
			zeroes = append(zeroes, row)
		} else {
			ones = append(ones, row)
		}
	}

	zeroLength := len(zeroes)
	oneLength := len(ones)
	if useHigher {
		if zeroLength == oneLength {
			return findNum(ones, charIndex+1, true)
		} else if zeroLength > oneLength {
			return findNum(zeroes, charIndex+1, true)
		} else {
			return findNum(ones, charIndex+1, true)
		}
	}

	if zeroLength == oneLength {
		return findNum(zeroes, charIndex+1, false)
	} else if zeroLength > oneLength {
		return findNum(ones, charIndex+1, false)
	} else {
		return findNum(zeroes, charIndex+1, false)
	}
}
