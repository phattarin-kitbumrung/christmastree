import express from 'express'

const app = express()
const port = 3000

app.listen(port, async () => {       
  console.log( `server started at http://localhost:${port}`)
  draw()
})

function draw(max = 4) {
  if (max < 1) {
    console.error("The number of levels must be at least 1.")
    return
  }

  const totalWidth = (max * 2) - 1 
  const center = Math.floor(totalWidth / 2) 

  for (let i = 0; i < max; i++) {
    const stars = '*'.repeat((i * 2) + 1) 
    const row = stars.padStart(center + i + 1, ' ').padEnd(totalWidth, ' ') 
    console.log(row)
  }

  // Print the vertical line below the pyramid
  const verticalLine = '|'.padStart(center + 1, ' ').padEnd(totalWidth, ' ')
  console.log(verticalLine)
}
