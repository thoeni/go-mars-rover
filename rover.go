package mars_rover

type Rover struct {
	x         int
	y         int
	direction string
}

type Grid struct {
	xMax int
	yMax int
}

var grid Grid

func RoverSetup() *Rover {
	var rover = new(Rover)
	rover.x = 0
	rover.y = 0
	rover.direction = "N"
	return rover
}

func InitGrid(xMax, yMax int) {
	grid.xMax = xMax
	grid.yMax = yMax
}

func (this *Rover) GetX() int {
	return this.x
}

func (this *Rover) GetY() int {
	return this.y
}

func (this *Rover) GetDirection() string {
	return this.direction
}

func (this *Rover) MoveForward() {
	switch this.GetDirection() {
	case "N":
		this.y = towardsNE(this.y, grid.yMax)
	case "S":
		this.y = towardsSW(this.y, grid.yMax)
	case "E":
		this.x = towardsNE(this.x, grid.xMax)
	case "W":
		this.x = towardsSW(this.x, grid.xMax)
	}
}

func towardsNE(current, max int) int {
	if current != max {
		return current + 1
	} else {
		return 0
	}
}

func towardsSW(current, max int) int {
	if current != 0 {
		return current - 1
	} else {
		return max
	}
}

func (this *Rover) MoveBackwards() {
	switch this.GetDirection() {
	case "N":
		this.y = towardsSW(this.y, grid.yMax)
	case "S":
		this.y = towardsNE(this.y, grid.yMax)
	case "E":
		this.x = towardsSW(this.x, grid.xMax)
	case "W":
		this.x = towardsNE(this.x, grid.xMax)
	}
}

func (this *Rover) TurnRight() {
	switch this.direction {
	case "N":
		this.direction = "E"
		break
	case "E":
		this.direction = "S"
		break
	case "S":
		this.direction = "W"
		break
	case "W":
		this.direction = "N"
		break
	}
}

func (this *Rover) TurnLeft() {
	this.TurnRight()
	this.TurnRight()
	this.TurnRight()
}
