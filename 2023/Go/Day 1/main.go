package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getDigits(input string) (int, error) {
	numString := ""
	for _, ch := range input {
		if '0' <= ch && ch <= '9' {
			numString += string(ch)
		}
	}
	log.Println("List of numbers:", numString)

	num, err := strconv.Atoi(string([]byte{numString[0], numString[len(numString)-1]}))

	if err != nil {
		log.Println("Error converting digits")
		return 0, err
	}

	log.Println("Calibration values: ", num)

	return num, nil
}

func main() {

	fmt.Print("Enter path: ")

	var path string
	_, err := fmt.Scan(&path)

	

	if err != nil {
		log.Fatalln("Error getting input path: \n", err)
	}
	log.Printf("Reading file from path: %s", path)

	dataBytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatalln("Error reading file:\n", err)
	}

	var dataString []string = strings.Split(string(dataBytes), "\n")
	var sum int = 0

	for _, input := range dataString {
		num, err := getDigits(input)
		if err != nil {
			return
		}
		sum += num
	}

	fmt.Println("\nSum of all of the Calibration values: ", sum)
}
