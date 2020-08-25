# Package for matrix math on Go slices

[![GoDev](https://img.shields.io/static/v1?label=godev&message=reference&color=00add8)][godev]
[![Build Status](https://travis-ci.org/google/go-cmp.svg?branch=master)][coverage]
This package is intended to be a tool for slice matrix math, specifically multiplication, addition, subtraction.

The primary features of `matrix` are:

* To perform matrix multiplication.

* To perform matrix addition.

* To perform matrix subtraction.

* To perform matrix transpose.

This is not an official Google product.

[godev]: https://github.com/randysimpson/go-matrix#matrix-methods
[coverage]: https://github.com/randysimpson/go-matrix/blob/master/coverage.html

## Install

```
go get -u github.com/randysimpson/go-matrix/matrix
```

## Usage
Go example of adding 2x2 and 2x2 matrix together:

```go
package main

import (
  "github.com/randysimpson/go-matrix/matrix"
  "fmt"
)

func main() {
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

  result := matrix.Add(a, b)
  
  fmt.Printf("result matrix: %v\n", result)
}
```

The resulting output is:

```sh
result matrix: [[3 5] [8 11]]
```

## Matrix methods

### func Add
```
Add(a [][]float64, b [][]float64) [][]float64
```
Add returns slice of slices to represent the sum of two matrices.

### func AddScalar
```
AddScalar(a [][]float64, b float64) [][]float64
```
AddScalar returns slice of slices to represent the sum of a matrix and a scalar.

### func Multiply
```
Multiply(a [][]float64, b [][]float64) [][]float64
```
Multiply returns slice of slices to represent the product of two matrices.

### func MultiplyScalar
```
MultiplyScalar(a [][]float64, b float64) [][]float64
```
MultiplyScaler returns slice of slices to represent the product of a matrix and a scalar.

### func Subtract
```
Subtract(a [][]float64, b [][]float64) [][]float64
```
Subtract returns slice of slices to represent two matrices being subtracted.

### func SubtractScalar
```
SubtractScalar(a [][]float64, b float64) [][]float64
```
SubtractScalar returns slice of slices to represent the a matrix having a scalar subtracted.

### func Transpose
```
Transpose(a [][]float64) [][]float64
```
Subtract returns slice of slices to represent the transpose of a matrix.

## License

MIT - See [LICENSE][license] file

[license]: https://github.com/randysimpson/go-matrix/blob/master/LICENCE

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
SOFTWARE.