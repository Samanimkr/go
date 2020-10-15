package main

func aliveNeighbours(world [][]byte, x int, y int) int {
	var a = 0

	var prex int = x - 1
	var prey int = y - 1
	var nextx int = x + 1
	var nexty int = y + 1

	if x == 0 {
		prex = 15
	}
	if y == 0 {
		prey = 15
	}
	if x == 15 {
		nextx = 0
	}
	if y == 15 {
		nexty = 0
	}

	if world[prey][prex] == 255 {
		a++
	}
	if world[y][prex] == 255 {
		a++
	}
	if world[nexty][prex] == 255 {
		a++
	}
	if world[prey][x] == 255 {
		a++
	}
	if world[nexty][x] == 255 {
		a++
	}
	if world[prey][nextx] == 255 {
		a++
	}
	if world[y][nextx] == 255 {
		a++
	}
	if world[nexty][nextx] == 255 {
		a++
	}

	return a
}

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := make([][]byte, p.imageHeight)
	for i := range newWorld {
		newWorld[i] = make([]byte, p.imageWidth)
	}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			var a = aliveNeighbours(world, x, y)

			if world[y][x] == 255 {
				if a == 2 || a == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			} else {
				if a == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0
				}
			}
		}
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	var cells []cell

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] == 255 {
				cells = append(cells, cell{x: x, y: y})
			}
		}
	}

	return cells
}
