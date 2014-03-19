package main

import (
	"fmt"
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
	//let's call that 10 for now
	var earth [10]int

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
	var dataset = []Outline{ Outline{1, 3, 3}, Outline{2, 5, 6}, Outline{4, 2, 5}, Outline{4, 6, 7} }

	return dataset
}