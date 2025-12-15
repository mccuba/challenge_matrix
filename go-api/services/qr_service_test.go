package services

import (
    "math"
    "testing"
)
//import "fmt"

func TestComputeQR(t *testing.T) {
    matrix := [][]float64{
        {1, 2},
        {3, 4},
        {5, 6},
    }

    Q, R, err := ComputeQR(matrix)
    //fmt.Println("Q:", Q)
    //fmt.Println("R:", R)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    if len(Q) != 3 {
        t.Errorf("expected Q rows=3, got %d", len(Q))
    }

    if len(R) != 3 {
        t.Errorf("expected R rows=2, got %d", len(R))
    }
}

const eps = 1e-6

func TestComputeQR_Reconstruction(t *testing.T) {
    A := [][]float64{
        {1, 2},
        {3, 4},
        {5, 6},
    }

    Q, R, err := ComputeQR(A)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }

    // check dimensions are compatible
    if len(Q) != len(A) {
        t.Fatalf("Q rows mismatch: got %d, want %d", len(Q), len(A))
    }

    if len(Q[0]) != len(R) {
        t.Fatalf("Q cols must match R rows")
    }

    if len(R[0]) != len(A[0]) {
        t.Fatalf("R cols mismatch: got %d, want %d", len(R[0]), len(A[0]))
    }

    // verify A ≈ Q·R
    Arec := multiply(Q, R)
    if !matricesAlmostEqual(A, Arec, eps) {
        t.Errorf("A != Q·R within tolerance")
    }
}




func multiply(A, B [][]float64) [][]float64 {
    rows := len(A)
    cols := len(B[0])
    mid := len(B)

    result := make([][]float64, rows)
    for i := 0; i < rows; i++ {
        result[i] = make([]float64, cols)
        for j := 0; j < cols; j++ {
            for k := 0; k < mid; k++ {
                result[i][j] += A[i][k] * B[k][j]
            }
        }
    }
    return result
}

func matricesAlmostEqual(A, B [][]float64, eps float64) bool {
    for i := range A {
        for j := range A[i] {
            if math.Abs(A[i][j]-B[i][j]) > eps {
                return false
            }
        }
    }
    return true
}

