package utils

func MiniMaxFloat(arr []float64) (float64, float64) {
	smallest, biggest := arr[0], arr[0]
	for _, v := range arr {
		if v > biggest {
			biggest = v
		}
		if v < smallest {
			smallest = v
		}
	}
	return smallest, biggest
}
