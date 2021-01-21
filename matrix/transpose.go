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

// transposeCol uses channel for queuing up result of the return matrix rows.
func transposeCol(m [][]float64, index int, ch chan []float64) {
	var result []float64
	result = append(result, float64(index))
	for r := 0; r < len(m); r++ {
		result = append(result, m[r][index])
	}
	ch <- result
}

// Transpose returns slice of slices to represent the transpose of a matrix.
func Transpose(m [][]float64) ([][]float64) {
	var result [][]float64
	ch := make(chan []float64)

	for c := 0; c < len(m[0]); c++ {
		go transposeCol(m, c, ch)
		result = append(result, []float64{})
	}
	for c := 0; c < len(m[0]); c++ {
		rtn := <-ch
		result[int(rtn[0])] = rtn[1:]
	}

	return result
}