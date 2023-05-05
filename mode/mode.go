package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func mode(file string) ([]int, error) {
	// open the file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// read the file line by line
	counts := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// convert the line to an integer
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		// count the integers
		counts[num]++
	}
	fmt.Println("number of occurances of each value from file: ", counts)

	// find the mode(s)
	var modeNums []int
	var ModeCount int
	for num, count := range counts {
		if count > ModeCount {
			modeNums = []int{num}
			ModeCount = count
		} else if count == ModeCount {
			modeNums = append(modeNums, num)
		}
	}
	fmt.Println("occurrences of mode value:", ModeCount)

	return modeNums, nil
}

func main() {
	modeNums, err := mode("dataM.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Print the mode value(s)
	fmt.Println("Mode(s) value:", modeNums)
	//fmt.Println("number of lines in dataM.txt:", lineCount)
	//fmt.Println("the modulus number: %d, had %d occurrences", modeNums, ModeCount)
}
