package main

/**
Inefficient brute force method of solving kangaroo
https://www.hackerrank.com/challenges/kangaroo
**/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//declare variables
	var x1, v1, x2, v2 int
	var check bool

	Scanner := bufio.NewScanner(os.Stdin)

	Scanner.Scan()
	scannedInput := Scanner.Text()
	sliceOfInput := make([]int, 4)
	for i := 0; i < 4; i++ {
		//fmt.Scanf("%v %v %v %v", &x1, &v1, &x2, &v2) might me more efficient
		sliceOfInput[i], _ = strconv.Atoi(strings.Split(scannedInput, " ")[i])
	}

	x1 = sliceOfInput[0]
	v1 = sliceOfInput[1]
	x2 = sliceOfInput[2]
	v2 = sliceOfInput[3]
	//One of the recursive break conditions is breaking once the kangaroo that was initially behind is ahead of the other kangaroo, hence order is important
	if x1 > x2 {
		fmt.Printf("%d>%d", x1, x2)
		check = existAtSamePointAtSameIteration(x1, v1, x2, v2)
	} else if x2 > x1 {
		fmt.Printf("%d>%d", x2, x1)
		check = existAtSamePointAtSameIteration(x2, v2, x1, v1)
	} else {
		check = true
	}

	if check {
		fmt.Printf("%s", "YES")
	} else {
		fmt.Printf("%s", "NO")
	}

}

//Wrapper function for timer function
func existAtSamePointAtSameIteration(x1, v1, x2, v2 int) bool {
	//Initiate timer to 0
	return timerFunction(x1, v1, x2, v2, 0)

}

func timerFunction(x1, v1, x2, v2, timer int) bool {
	incrementedx1 := x1 + v1
	incrementedx2 := x2 + v2
	if incrementedx1 == incrementedx2 {
		return true
	} else if incrementedx1 < incrementedx2 {
		return false
	} else if timer > 100000 {
		return false
	} else {
		return timerFunction(incrementedx1, v1, incrementedx2, v2, timer+1)
	}
}

func convertToIntFromASCII(x int) int {
	return (x - 48)
}
