package skylinecommon

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Outline struct {
	L int
	H int
	R int
}

func GetTestData() Outline[] {
	file, _ := os.Open("skylineinput.txt")
	defer file.Close()

	var outlines []Outline
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit :=  strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(lineSplit[0])
		h, _ := strconv.Atoi(lineSplit[1])
		y, _ := strconv.Atoi(lineSplit[2])

		outlines = append(outlines, Outline { x, h, y })
	}

	//var dataset = []Outline{ Outline{1, 11, 5}, Outline{2, 6, 7}, Outline{3, 13, 9}, Outline{12, 7, 6}, Outline{14, 3, 25}, Outline{19, 18, 22}, Outline{23, 13, 29}, Outline{24, 4, 28} }

	return outlines
}