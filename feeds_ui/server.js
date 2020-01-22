/* eslint-disable-next-line @typescript-eslint/no-var-requires */
const express = require('express')
/* eslint-disable-next-line @typescript-eslint/no-var-requires */
const path = require('path')

const port = process.env.PORT || 8080
const host = '0.0.0.0'

const app = express()

app.use(express.static(__dirname))
app.use(express.static(path.join(__dirname, 'build')))

app.get('/*', function(req, res) {
  res.sendFile(path.join(__dirname, 'build', 'index.html'))
})

app.listen(port, host, () => {
  console.log(`Server started... on ${host}:${port}`)
})
