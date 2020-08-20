package matrix

func Subtract(m1 [][]float64, m2 [][]float64) ([][]float64) {
	//check to see if m1 and m2 have the same dimensions
	var result [][]float64

	for r := 0; r < len(m1); r++ {
		for c := 0; c < len(m1[r]); c++ {
			val := m1[r][c] - m2[r][c]

			//put into results
			if len(result) - 1 < r {
				var inner = []float64{val}
				result = append(result, inner)
			} else {
				//need to add to inner results.
				result[r] = append(result[r], val)
			}
		}
	}

	return result
}