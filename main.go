package mars_rover

import "strings"

func CommandRover(rover *Rover, command string) {
	for _, v := range command {
		switch strings.ToUpper(string(v)) {
		case "F":
			rover.MoveForward()
			break
		case "B":
			rover.MoveBackwards()
			break
		case "R":
			rover.TurnRight()
			break
		case "L":
			rover.TurnLeft()
			break
		}
	}
}
