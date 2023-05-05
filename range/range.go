package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func rangeValue(file string) (int, error) {
	// open the file
	f, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// read the file line by line
	var nums []int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// convert the line to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			return 0, err
		}
		nums = append(nums, num)
	}

	// calculate the range
	if len(nums) == 0 {
		return 0, nil
	}
	min := nums[0]
	max := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	rangeValue := max - min

	return rangeValue, nil
}

func main() {
	rangeValue, err := rangeValue("dataR.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Range value is: ", rangeValue)
}
