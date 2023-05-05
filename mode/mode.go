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
		// count the integer
		counts[num]++
	}

	// find the mode(s)
	var modeNums []int
	var modeCount int
	for num, count := range counts {
		if count > modeCount {
			modeNums = []int{num}
			modeCount = count
		} else if count == modeCount {
			modeNums = append(modeNums, num)
		}
	}

	return modeNums, nil
}

func main() {
	modeNums, err := mode("dataM.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mode(s) value:", modeNums)
	//fmt.Println("number of lines in dataM.txt:", lineCount)
}
