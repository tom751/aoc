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

	var nums []int
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var sums []int
	for i := range nums {
		if i >= (len(nums) - 2) {
			break
		}

		indexTo := i + 3
		numsToAdd := nums[i:indexTo]

		sum := 0
		for _, num := range numsToAdd {
			sum += num
		}

		sums = append(sums, sum)
	}

	var prevNum int
	index := 0
	counter := 0
	for _, num := range sums {
		if index > 0 {
			if num > prevNum {
				counter++
			}
		}

		index++
		prevNum = num
	}

	fmt.Println(counter)
}
