package matrix_test

import (
	"github.com/randysimpson/go-matrix/matrix"
	//"go-matrix/matrix"
	"testing"
	"reflect"
)

func TestAdd(t *testing.T) {
	var ansA = []float64{3.0, 5.0}
	var ansB = []float64{8.0, 11.0}
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

	if got := matrix.Add(a, b); !reflect.DeepEqual(got, want) {
			t.Errorf("Add(a, b) = %v, want %v", got, want)
	}
}
