const express = require("express")
const router = express.Router()
const { computeStats } = require("../services/statsService")

router.post("/", (req, res) => {
  const { Q, R } = req.body
  const stats = computeStats(Q, R)
  res.json(stats)
})

module.exports = router

