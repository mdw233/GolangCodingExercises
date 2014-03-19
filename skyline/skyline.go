package main

import (
	"fmt"
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


//problem here: https://stackoverflow.com/questions/1066234/the-skyline-problem
//and here: http://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=3&page=show_problem&problem=41
func main() {

	//get a list of points
	var buildings = getTestDataSet()

	//build up an int array large enough to handle all points
	//let's call that 40 for now
	var earth [40]int

	//iterate over the list of outlines
	// check the left border - is the height larger than previous heights - if so, record the new height
	// do the same thing for each element between l and r
	for _, o := range buildings {	
		for j:= o.L; j <= o.R; j++ {
			if (earth[j] < o.H) {
				earth[j] = o.H
			}
		}
	}

	for i, p := range earth {
		fmt.Println(i, p)
	}

}

func getTestDataSet() []Outline {

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
