package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func mode(file string) (int, error) {
	// open the file
	file, err := os.Open("data.txt")
	if err != nil {
		return 0, err
	}
	defer f.Close()

	// read the file line by line
	counts := make(map[int]int)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		// split the line by commas
		parts := strings.Split(line, ",")
		// convert each part to an integer and count it
		for _, part := range parts {
			num, err := strconv.Atoi(part)
			if err != nil {
				return 0, err
			}
			counts[num]++
		}
	}

	// find the mode
	var modeNum int
	var modeCount int
	for num, count := range counts {
		if count > modeCount {
			modeNum = num
			modeCount = count
		}
	}

	return modeNum, nil
}

func main() {
	fmt.Println("Welcome to the Math Skills Project")

	// Open the external data file for reading
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("error reading external file")
		panic(err)
	}
	modeNum, err := mode("data.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println("Mode:", modeNum)
	defer file.Close()

	// Initialize variables for calculating the average mean
	sum := 0
	count := 0

	// Read (scan) values from the external file, convert them to integars, and calculate the sum and count
	var values []int
	var valueStr string
	for {
		_, err := fmt.Fscanln(file, &valueStr)
		if err != nil {
			break
			// EoF
		}
		// Convert the string from the external file to an integar
		num, err := strconv.Atoi(valueStr)
		if err != nil {
			panic(err)
		}
		sum += num
		count++
		values = append(values, num)
	}
	// calculate the average mean and print the result to the screen
	aveMean := float64(sum) / float64(count)
	fmt.Printf("average mean of values: %2f\n", aveMean)

	// sort the values from the file in ascending order
	sort.Ints(values)

	// calculate the median (middle position) of the list of values
	var median float64
	countm := len(values)
	if countm%2 == 0 {
		// average the middle two values if there is an even number of values
		middle1 := values[countm/2-1]
		middle2 := values[countm/2]
		median = float64(middle1+middle2) / 2.0
	} else {
		// take the middle value if there are an odd number of values
		middle := values[countm/2]
		median = float64(middle)
	}
	// print the median (middle) value from the list
	fmt.Printf("Median (middle) is: %.2f\n", median)

	// Calculate the Average Mean of the values from the external file for Variance
	meanv := 0.0
	countv := float64(len(values))
	for _, numv := range values {
		meanv += float64(numv) / countv
	}
	// Calculate the Variance of the Average Mean value from the external file
	variancev := 0.0
	for _, numv := range values {
		diff := float64(numv) - meanv
		variancev += math.Pow(diff, 2) / countv
	}
	// print the Variance value
	fmt.Printf("Variance value: %.2f\n", variancev)

	// Calculate the Standard Deviation of the values from the external file via the Variance value
	standardDeviation := math.Sqrt(variancev)
	// print the Standard Deviation value
	fmt.Printf("Standard Deviation value: %.2f\n", standardDeviation)

	var modeNum int
	var modeCount int
	for nump, count := range counts {
		if count > modeCount {
			modeNum = nump
			modeCount = count
		}
	}
	fmt.Println("Mode:", modeNum)
}

// Calculate the Mode of the values from the external file
func mode(file string) (int, error) {
	// open the external file
	file, err := os.Open(file)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	// read the external file line by line
	counts := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// split the line by commas
		parts := strings.Split(line, ",")
		// convert each part to an integer and count the number of parts
		for _, part := range parts {
			nump, err := strconv.Atoi(part)
			if err != nil {
				return
			}
			counts[nump]++
		}
	}
}
