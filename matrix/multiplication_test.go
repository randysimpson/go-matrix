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

//test 2 x 3 times 3 x 2 Multiply
func TestMultiply2x3_3x2(t *testing.T) {
	var ansA = []float64{58.0, 64.0}
	var ansB = []float64{139.0, 154.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)

	var testA = []float64{1.0, 2.0, 3.0}
	var testB = []float64{4.0, 5.0, 6.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)

	var testC = []float64{7.0, 8.0}
	var testD = []float64{9.0, 10.0}
	var testE = []float64{11.0, 12.0}
	var b [][]float64
	b = append(b, testC)
	b = append(b, testD)
	b = append(b, testE)

	if got := matrix.Multiply(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Multiply(a, b) = %v, want %v", got, want)
	}
}

//test 1 x 3 times 3 x 4 Multiply
func TestMultiply1x3_3x4(t *testing.T) {
	var ansA = []float64{83.0, 63.0, 37.0, 75.0}
	var want [][]float64
	want = append(want, ansA)

	var testA = []float64{3.0, 4.0, 2.0}
	var a [][]float64
	a = append(a, testA)

	var testB = []float64{13.0, 9.0, 7.0, 15.0}
	var testC = []float64{8.0, 7.0, 4.0, 6.0}
	var testD = []float64{6.0, 4.0, 0.0, 3.0}
	var b [][]float64
	b = append(b, testB)
	b = append(b, testC)
	b = append(b, testD)

	if got := matrix.Multiply(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Multiply(a, b) = %v, want %v", got, want)
	}
}

func TestMultiplyScalar(t *testing.T) {
	var ansA = []float64{5.0, 10.0}
	var ansB = []float64{15.0, 20.0}
	var want [][]float64
	want = append(want, ansA)
	want = append(want, ansB)

	var testA = []float64{1.0, 2.0}
	var testB = []float64{3.0, 4.0}
	var a [][]float64
	a = append(a, testA)
	a = append(a, testB)

	if got := matrix.MultiplyScalar(a, 5.0); !reflect.DeepEqual(got, want) {
		t.Errorf("MultiplyScaler(a, b) = %v, want %v", got, want)
	}
}
