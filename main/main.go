package main

import (
	"fmt"
	"race_condition_in_go/fix"
	"race_condition_in_go/problem"
)

func main() {
	problem.ProblemDemo()
	fix.FixUsingMutex()
	fix.FixUsingChannel()
	fmt.Println("Sagar")
}
