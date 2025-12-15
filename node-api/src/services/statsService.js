const { flatten, isDiagonal } = require("../utils/matrixUtils")

function computeStats(Q, R) {
  const values = flatten([Q, R])

  const sum = values.reduce((a, b) => a + b, 0)

  return {
    max: Math.max(...values),
    min: Math.min(...values),
    sum,
    avg: sum / values.length,
    hasDiagonal: isDiagonal(Q) || isDiagonal(R),
  }
}

module.exports = { computeStats }

