package main

import (
	"fmt"
	"./common"
)


//problem here: https://stackoverflow.com/questions/1066234/the-skyline-problem
//and here: http://uva.onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=3&page=show_problem&problem=41
func main() {

	//get a list of points
	var buildings = skylinecommon.GetTestData()

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

	//1. now draw the line
	//2. find the first non 0 height - that index is the starting point
	//	find the height of that - that's the first x
	//3. continue over the array - each time the height changes, record the new index, then height
	//4. at the very end put a 0

	//may actually be able to short cicuit all of this and just start at 3.

	for i, p := range earth {
		fmt.Println(i, p)
	}
}
