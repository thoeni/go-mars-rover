package marsrover

type rover struct {
	x         int
	y         int
	direction string
}

type gridPlanet struct {
	xMax int
	yMax int
}

var grid gridPlanet

const n string = "N"
const e string = "E"
const s string = "S"
const w string = "W"

func roverSetup() *rover {
	var rover = new(rover)
	rover.x = 0
	rover.y = 0
	rover.direction = n
	return rover
}

func initGrid(xMax, yMax int) {
	grid.xMax = xMax
	grid.yMax = yMax
}

func (rover *rover) getX() int {
	return rover.x
}

func (rover *rover) getY() int {
	return rover.y
}

func (rover *rover) getDirection() string {
	return rover.direction
}

func (rover *rover) moveForward() {
	switch rover.getDirection() {
	case n:
		rover.y = towardsNE(rover.y, grid.yMax)
	case s:
		rover.y = towardsSW(rover.y, grid.yMax)
	case e:
		rover.x = towardsNE(rover.x, grid.xMax)
	case w:
		rover.x = towardsSW(rover.x, grid.xMax)
	}
}

func towardsNE(current, max int) int {
	if current != max {
		return current + 1
	}
	return 0
}

func towardsSW(current, max int) int {
	if current != 0 {
		return current - 1
	}
	return max
}

func (rover *rover) moveBackwards() {
	switch rover.getDirection() {
	case n:
		rover.y = towardsSW(rover.y, grid.yMax)
	case s:
		rover.y = towardsNE(rover.y, grid.yMax)
	case e:
		rover.x = towardsSW(rover.x, grid.xMax)
	case w:
		rover.x = towardsNE(rover.x, grid.xMax)
	}
}

func (rover *rover) turnRight() {
	switch rover.direction {
	case n:
		rover.direction = e
		break
	case e:
		rover.direction = s
		break
	case s:
		rover.direction = w
		break
	case w:
		rover.direction = n
		break
	}
}

func (rover *rover) turnLeft() {
	rover.turnRight()
	rover.turnRight()
	rover.turnRight()
}
