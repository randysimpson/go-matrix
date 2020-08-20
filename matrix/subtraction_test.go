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

//test 2 x 2 minus 2 x 2
func TestSubtract(t *testing.T) {
	var ansA = []float64{-1.0, -1.0}
	var ansB = []float64{-2.0, -3.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)

	var testA = []float64{1.0, 2.0}
	var testB = []float64{3.0, 4.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)

	var testC = []float64{2.0, 3.0}
	var testD = []float64{5.0, 7.0}
	var b [][]float64
	b = append(b, testC)
	b = append(b, testD)

	if got := matrix.Subtract(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Subtract(a, b) = %v, want %v", got, want)
	}
}

func TestSubtractScalar(t *testing.T) {
	var ansA = []float64{-4.0, -3.0}
	var ansB = []float64{-2.0, -1.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)

	var testA = []float64{1.0, 2.0}
	var testB = []float64{3.0, 4.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)

	if got := matrix.SubtractScalar(a, 5.0); !reflect.DeepEqual(got, want) {
		t.Errorf("SubtractScalar(a, b) = %v, want %v", got, want)
	}
}
