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

// Subtract returns slice of slices to represent the transpose of a matrix.
func Transpose(m1 [][]float64) ([][]float64) {
	var result [][]float64

	for r := 0; r < len(m1); r++ {
		for c := 0; c < len(m1[r]); c++ {
			//put into results
			if len(result) - 1 < c {
				var inner = []float64{m1[r][c]}
				result = append(result, inner)
			} else {
				//need to add to inner results.
				result[c] = append(result[c], m1[r][c])
			}
		}
	}

	return result
}