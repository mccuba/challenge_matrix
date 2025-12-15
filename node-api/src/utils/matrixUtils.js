module.exports = { flatten, isDiagonal }

function validateMatrix(matrix) {
  if (!Array.isArray(matrix) || matrix.length === 0) {
    throw new Error("Matrix must be a non-empty array")
  }

  const cols = matrix[0].length

  if (!Array.isArray(matrix[0]) || cols === 0) {
    throw new Error("Matrix rows must be non-empty arrays")
  }

  for (const row of matrix) {
    if (!Array.isArray(row) || row.length !== cols) {
      throw new Error("Matrix is not rectangular")
    }

    for (const value of row) {
      if (typeof value !== "number" || Number.isNaN(value)) {
        throw new Error("Matrix contains non-numeric values")
      }
    }
  }
}

function isDiagonal(matrix) {
  if (matrix.length !== matrix[0].length) return false
  for (let i = 0; i < matrix.length; i++) {
    for (let j = 0; j < matrix.length; j++) {
      if (i !== j && matrix[i][j] !== 0) return false
    }
  }
  return true
}


function flatten(matrices) {
  const values = []

  for (const matrix of matrices) {
    validateMatrix(matrix)

    for (const row of matrix) {
      for (const value of row) {
        values.push(value)
      }
    }
  }

  return values
}

