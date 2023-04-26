package pcl

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type DType string

var (
	ASCII             DType = "ascii"
	BINARY            DType = "binary"
	BINARY_COMPRESSED DType = "binary_compressed"
)

type PCD struct {
	HEADER struct {
		RawText string
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
		DATA      DType
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

	p := new(PCD)

	// handle header
	for scanner.Scan() {
		// if the text starts with "DATA ", then break
		if strings.HasPrefix(scanner.Text(), "DATA ") {
			p.updateHeader(scanner.Text())
			p.HEADER.RawText += "DATA "
			break
		}

		p.updateHeader(scanner.Text())
		p.HEADER.RawText += scanner.Text() + "\n"

	}

	// handle the data
	scanner.Scan()
	//todo
	p.updateData(scanner)
	log.Println(p)

	return nil

}

func (p *PCD) updateHeader(line string) {
	// split the line
	words := strings.Split(line, " ")
	switch words[0] {
	case "VERSION":
		p.HEADER.VERSION = words[1]
	case "FIELDS":
		p.HEADER.FIELDS = make([]struct {
			NAME string
			TYPE string
			SIZE int
		}, len(words)-1)
		for i := 1; i < len(words); i++ {
			p.HEADER.FIELDS[i-1].NAME = words[i]
		}

	case "SIZE":
		if p.HEADER.FIELDS == nil {
			p.HEADER.FIELDS = make([]struct {
				NAME string
				TYPE string
				SIZE int
			}, len(words)-1)

			for i := 1; i < len(words); i++ {
				p.HEADER.FIELDS[i-1].SIZE, _ = strconv.Atoi(words[i])
			}

		} else {
			for i := 1; i < len(words); i++ {
				p.HEADER.FIELDS[i-1].SIZE, _ = strconv.Atoi(words[i])
			}
		}
	case "TYPE":
		if p.HEADER.FIELDS == nil {
			p.HEADER.FIELDS = make([]struct {
				NAME string
				TYPE string
				SIZE int
			}, len(words)-1)

			for i := 1; i < len(words); i++ {
				p.HEADER.FIELDS[i-1].TYPE = words[i]
			}

		} else {
			for i := 1; i < len(words); i++ {
				p.HEADER.FIELDS[i-1].TYPE = words[i]
			}
		}
	case "WIDTH":
		p.HEADER.WIDTH, _ = strconv.Atoi(words[1])
	case "HEIGHT":
		p.HEADER.HEIGHT, _ = strconv.Atoi(words[1])
	case "VIEWPOINT":
		p.HEADER.VIEWPOINT = make([]float64, len(words)-1)
	case "POINTS":
		p.HEADER.POINTS, _ = strconv.Atoi(words[1])
	case "DATA":
		switch words[1] {
		case "ascii":
			p.HEADER.DATA = ASCII
		case "binary":
			p.HEADER.DATA = BINARY
		case "binary_compressed":
			p.HEADER.DATA = BINARY_COMPRESSED
		default:
			panic("DATA type not supported")
		}
	}
}

func (p *PCD) updateData(s *bufio.Scanner) {
	pointSize := 0
	for _, field := range p.HEADER.FIELDS {
		pointSize += field.SIZE
	}
	if p.HEADER.DATA == ASCII {
		// todo
	} else if p.HEADER.DATA == BINARY {

		d := s.Bytes()
		log.Println(d)

	} else if p.HEADER.DATA == BINARY_COMPRESSED {
		// todo
	} else {
		panic("DATA type not supported")

	}
}
