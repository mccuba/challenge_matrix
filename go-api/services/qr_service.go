package services

import (
    "errors"
    "gonum.org/v1/gonum/mat"
)

func ComputeQR(matrix [][]float64) ([][]float64, [][]float64, error) {
    if len(matrix) == 0 {
        return nil, nil, errors.New("empty matrix")
    }

    rows := len(matrix)
    cols := len(matrix[0])

    data := make([]float64, 0, rows*cols)
    for _, row := range matrix {
        if len(row) != cols {
            return nil, nil, errors.New("non-rectangular matrix")
        }
        data = append(data, row...)
    }

    m := mat.NewDense(rows, cols, data)

    var qr mat.QR
    qr.Factorize(m)

    var q, r mat.Dense
    qr.QTo(&q)
    qr.RTo(&r)

    return denseToSlice(&q), denseToSlice(&r), nil
}

func denseToSlice(d *mat.Dense) [][]float64 {
    r, c := d.Dims()
    out := make([][]float64, r)
    for i := 0; i < r; i++ {
        out[i] = make([]float64, c)
        for j := 0; j < c; j++ {
            out[i][j] = d.At(i, j)
        }
    }
    return out
}

