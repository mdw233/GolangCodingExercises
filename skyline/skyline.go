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

	//iterate over array
	//1. start by finding the first non-zero height
	//2. when the height at an index changes, mark the difference between this and the last index change and add that to the line array
	//3. then add the height to the line array
	//4. unless the new height is 0, then mark that height of the in
	var skyline []int
	indexOfLastChange, lastHeight := 0, 0
	for i := 0; i < len(earth); i++ {
		if earth[i] == 0 {
			if lastHeight > 0 {
				skyline = append(skyline, i - 1)
				skyline = append(skyline, 0)

				indexOfLastChange = i
				lastHeight = earth[i]
			}
		} else {
			if lastHeight == 0 {
				skyline = append(skyline, i)
				skyline = append(skyline, earth[i])

				indexOfLastChange = i
				lastHeight = earth[i]
			} else if lastHeight != earth[i] {
				//draw the horizontal line
				skyline = append(skyline, i - indexOfLastChange)
				skyline = append(skyline, earth[i])

				indexOfLastChange = i
				lastHeight = earth[i]
			}
		}
	}

	//may actually be able to short cicuit all of this and just start at 3.

	for i := 0; i < len(skyline); i++ {
		fmt.Println(skyline[i])
	}
}
