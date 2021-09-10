package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strings"

	exif "github.com/xor-gate/goexif2/exif"
)

func main() {
	fmt.Println("Starting program...")

	var filePath string

	// minimum values for location offset
	min := 0.00000000000
	// get the path from the comman line argument
	if len(os.Args) == 2 {
		filePath = os.Args[1]
	}

	// if command line argument is not present, ask the user for path
	if len(os.Args) < 2 {
		fmt.Print("Path to your image: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')

		// remove new line from the string
		input = strings.TrimSuffix(input, "\n")
		fmt.Println(input)
		filePath = input
	}

	fmt.Println("file: ", filePath)

	// Get coordinates of the image
	lat, lon, err := getPosition(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("original lat: ", lat)
	fmt.Println("original lon: ", lon)
	lat = lat + min + rand.Float64()
	lon = lon + min + rand.Float64()
	fmt.Println("\nchanged lat: ", lat)
	fmt.Println("changed long: ", lon)

	// Change the coordinates of your photo
	changePosition(filePath, lat, lon)
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

func changePosition(filePath string, lat, lon float64) {
	var out bytes.Buffer

	cmd := exec.Command(
		"exiftool", filePath, fmt.Sprintf("-gpslatitude=%f", lat), fmt.Sprintf("-gpslongitude=%f", lon))
	cmd.Stdout = &out
	cmd.Stderr = &out

	err := cmd.Run()
	if err != nil {
		panic(out.String())
	}
	fmt.Println("\nlocation changed successfuly!")
}
