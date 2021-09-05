package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	exif "github.com/xor-gate/goexif2/exif"
)

func main() {
	fmt.Println("Starting program...")

	var filePath string

	// if command line argument is not present, ask the user for path
	if len(os.Args) <= 2 {
		fmt.Print("Path to your image: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		// remove new line from the string
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)
		filePath = input
	}

	// get the path from the comman line argument
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}
	fmt.Println("file: ", filePath)

	// Get coordinate of the image
	lat, lon, err := getPosition(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("lat: ", lat)
	fmt.Println("lon: ", lon)
}

func getPosition(filePath string) (float64, float64, error) {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}

	fileDecoded, err := exif.Decode(file)
	if err != nil {
		return 0, 0, err
	}

	lat, lon, err := fileDecoded.LatLong()
	if err != nil {
		return 0, 0, err
	}

	return lat, lon, nil
}
