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
	file, _ := os.Open("skylineinput.txt")
	defer file.Close()

	var buildings []Outline
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplit :=  strings.Split(scanner.Text(), " ")
		x, _ := strconv.Atoi(lineSplit[0])
		h, _ := strconv.Atoi(lineSplit[1])
		y, _ := strconv.Atoi(lineSplit[2])

		buildings = append(buildings, Outline { x, h, y })
	}

	PrintSkyline(buildings)
}

func PrintSkyline(buildings []Outline) {
	
	//build up an int array large enough to handle all points
	//let's call that 5000 for now
	var earth [5000]int

	//iterate over the list of outlines
	// check the left border - is the height larger than previous heights - if so, record the new height
	// do the same thing for each element between l and r
	for _, o := range buildings {	
		for j:= o.L; j < o.R; j++ {
			if (earth[j] < o.H) {
				earth[j] = o.H
			}
		}
	}

	//1. iterate over array
	//2. when the height at an index changes, mark the difference between this and the last index change and add that to the line array
	//3. then add the height to the line array
	var skyline []int
	lastHeight := 0
	for i := 0; i < len(earth); i++ {
		if (lastHeight != earth[i]) {
			skyline = append(skyline, i)
			skyline = append(skyline, earth[i])
			lastHeight = earth[i]
		}
	}

	for i := 0; i < len(skyline); i++ {
		fmt.Println(skyline[i])
	}
}
