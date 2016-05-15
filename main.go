package marsrover

import "strings"

func commandRover(rover *rover, command string) {
	for _, v := range command {
		switch strings.ToUpper(string(v)) {
		case "F":
			rover.moveForward()
			break
		case "B":
			rover.moveBackwards()
			break
		case "R":
			rover.turnRight()
			break
		case "L":
			rover.turnLeft()
			break
		}
	}
}
