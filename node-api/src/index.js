//const express = require("express")
//const app = express()

//app.use(express.json())

//app.post("/stats", (req, res) => {
  //res.json({ message: "Stats endpoint working" })
//})

//app.listen(3000, () => {
  //console.log("Node API running on port 3000")
//})


const express = require("express")
const statsRoutes = require("./routes/stats")

const app = express()
app.use(express.json())

app.use("/stats", statsRoutes)

app.listen(3000, () => console.log("Node API running on 3000"))
