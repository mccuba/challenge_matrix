package models

type MatrixRequest struct {
    Matrix [][]float64 `json:"matrix"`
}

type QRResponse struct {
    Q [][]float64 `json:"Q"`
    R [][]float64 `json:"R"`
}
