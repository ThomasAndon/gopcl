package pcl

import (
	"bufio"
	"fmt"
	"os"
)

type PCD struct {
	HEADER struct {
		VERSION string
		FIELDS  []struct {
			NAME string
			TYPE string
			SIZE int
		}
		WIDTH     int
		HEIGHT    int
		VIEWPOINT []float64
		POINTS    int
	}
	DATA []Point
}

type Point struct {
	X         float64
	Y         float64
	Z         float64
	RGB       float64
	Intensity float64
}

func LoadPCDFile(path string) *PCD {
	// load file
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	// for read each line
	scanner := bufio.NewScanner(file)
	// for read each line
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil

}

func Hello() {
	println("Hello, v1.0.2")
}
