/*
MIT License

Copyright (c) 2020 Randy Simpson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.*/

package matrix

// multiplyRow uses channel for queuing up result of the return matrix rows.
func multiplyRow(s []float64, s2 [][]float64, index int, ch chan []float64) {
	var result []float64
	result = append(result, float64(index))
	for c := 0; c < len(s2[0]); c++ {
		val := 0.0
		for r := 0; r < len(s); r++ {
			val += s[r] * s2[r][c]
		}
		result = append(result, val)
	}
	ch <- result
}

// Multiply returns slice of slices to represent the product of two matrices.
func Multiply(m1 [][]float64, m2 [][]float64) ([][]float64) {
	//check to see if these can mulitply successfull by shape or throw error.
	var result [][]float64
	ch := make(chan []float64)

	for r := 0; r < len(m1); r++ {
		go multiplyRow(m1[r], m2, r, ch)
		result = append(result, []float64{})
	}
	for r := 0; r < len(m1); r++ {
		rtn := <-ch
		result[int(rtn[0])] = rtn[1:]
	}

	return result
}

// MultiplyScaler returns slice of slices to represent the product of a matrix and a scalar.
func MultiplyScalar(m [][]float64, num float64) ([][]float64) {
	var result [][]float64

	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[r]); c++ {
			val := m[r][c] * num

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
