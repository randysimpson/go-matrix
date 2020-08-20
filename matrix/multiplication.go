package matrix

func Multiply(m1 [][]float64, m2 [][]float64) ([][]float64) {
	//check to see if these can mulitply successfull by shape or throw error.
	var result [][]float64

	for r := 0; r < len(m1); r++ {
		for c := 0; c < len(m2[0]); c++ {
			val := 0.0
			for m1c := 0; m1c < len(m1[r]); m1c++ {
				val += m1[r][m1c] * m2[m1c][c]
			}

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