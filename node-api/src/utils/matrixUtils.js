function flatten(matrices) {
  return matrices.flat(2)
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

module.exports = { flatten, isDiagonal }

