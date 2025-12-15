const {
    computeStats
} = require("./statsService")
const {
    flatten
} = require("../utils/matrixUtils")

describe("computeStats", () => {
    test("throws error if matrix is not rectangular", () => {
        const matrices = [
            [
                [1, 2],
                [3], // fila inválida
            ],
        ]

        expect(() => flatten(matrices)).toThrow("Matrix is not rectangular")
    })
    test("flatten throws error if matrix contains non-numeric values", () => {
        const matrices = [
            [
                [1, 2],
                ["3", 4]
            ],
        ]

        expect(() => flatten(matrices)).toThrow(
            "Matrix contains non-numeric values"
        )
    })

    test("flatten throws error if matrix contains non-numeric values", () => {
        const matrices = [
            [
                [1, 2],
                [null, 4]
            ],
        ]

        expect(() => flatten(matrices)).toThrow(
            "Matrix contains non-numeric values"
        )
    })

    test("calculates correct statistics and detects diagonal matrix", () => {
        const Q = [
            [1, 0],
            [0, 1],
        ]

        const R = [
            [2, 3],
            [0, 4],
        ]

        const stats = computeStats(Q, R)

        expect(stats.max).toBe(4)
        expect(stats.min).toBe(0)
        expect(stats.sum).toBe(11)
        expect(stats.avg).toBeCloseTo(11 / 8)
        expect(stats.hasDiagonal).toBe(true)
    })

    test("detects when no matrix is diagonal", () => {
        const Q = [
            [1, 2],
            [3, 4],
        ]

        const R = [
            [5, 6],
            [7, 8],
        ]

        const stats = computeStats(Q, R)

        expect(stats.hasDiagonal).toBe(false)
    })
})
