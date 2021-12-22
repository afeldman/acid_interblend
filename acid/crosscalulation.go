package acid

import "math"

func AcidCrossCalculation(solution1, solution2, result_solution, volume float64) (float64, float64) {
	sol1 := math.Abs(result_solution - solution2)
	sol2 := math.Abs(result_solution - solution1)

	output_volume_s1 := volume * sol1 / (sol1 + sol2)
	output_volume_s2 := volume * sol2 / (sol1 + sol2)

	return output_volume_s1, output_volume_s2
}
