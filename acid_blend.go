package main

import (
	"fmt"
	"math"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	solution1       = kingpin.Arg("solution1", "strength of solution 1 in percent").Required().Float64()
	result_solution = kingpin.Arg("result_solution", "strength of solution for the end solution in percent").Required().Float64()
	solution2       = kingpin.Arg("solution2", "strength of solution 2 in percent").Default("0.0").Float64()
	volume          = kingpin.Flag("volume", "volume in mililiter").Default("1000.0").Float64()
)

func main() {
	kingpin.Parse()

	sol1 := math.Abs(*result_solution - *solution2)
	sol2 := math.Abs(*result_solution - *solution1)

	output_volume_s1 := *volume * sol1 / (sol1 + sol2)
	output_volume_s2 := *volume * sol2 / (sol1 + sol2)

	fmt.Println("Solution 1: ", fmt.Sprintf("%.2f", output_volume_s1), " ml")
	fmt.Println("Solution 2: ", fmt.Sprintf("%.2f", output_volume_s2), " ml")
}
