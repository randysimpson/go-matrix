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

package matrix_test

import (
	"github.com/randysimpson/go-matrix/matrix"
	"testing"
	"reflect"
)

//test 3 x 3
func TestTranspose3x3(t *testing.T) {
	var ansA = []float64{5.0, 4.0, 7.0}
	var ansB = []float64{4.0, 0.0, 10.0}
	var ansC = []float64{3.0, 4.0, 3.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)
	want = append(want, ansC)

	var testA = []float64{5.0, 4.0, 3.0}
	var testB = []float64{4.0, 0.0, 4.0}
	var testC = []float64{7.0, 10.0, 3.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)
	a = append(a, testC)

	if got := matrix.Transpose(a); !reflect.DeepEqual(got, want) {
			t.Errorf("Transpose(a) = %v, want %v", got, want)
	}
}

//test 4 x 2
func TestTranspose4x2(t *testing.T) {
	var ansA = []float64{5.0, 4.0, 7.0, -1.0}
	var ansB = []float64{4.0, 0.0, 10.0, 8.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)

	var testA = []float64{5.0, 4.0}
	var testB = []float64{4.0, 0.0}
	var testC = []float64{7.0, 10.0}
	var testD = []float64{-1.0, 8.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)
	a = append(a, testC)
	a = append(a, testD)

	if got := matrix.Transpose(a); !reflect.DeepEqual(got, want) {
			t.Errorf("Transpose(a) = %v, want %v", got, want)
	}
}