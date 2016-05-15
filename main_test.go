package marsrover

import "testing"

func TestCommandRoverFFRFF(t *testing.T) {
	t.Parallel()
	//Given
	expectedX := 2
	expectedY := 2
	expectedDirection := "E"
	initGrid(99, 99)
	rover := roverSetup()

	//When
	commandRover(rover, "ffrff")

	//Then
	actualX := rover.getX()

	if actualX != expectedX {
		t.Errorf("[X] Expected %d, actual %d", expectedX, actualX)
	}

	actualY := rover.getY()

	if actualY != expectedY {
		t.Errorf("[Y] Expected %d, actual %d", expectedY, actualY)
	}

	actualDirection := rover.getDirection()

	if actualDirection != expectedDirection {
		t.Errorf("[Direction] Expected %s, actual %s", expectedDirection, actualDirection)
	}
}

func TestCommandRoverBBLBB(t *testing.T) {
	t.Parallel()
	//Given
	expectedX := 2
	expectedY := 98
	expectedDirection := "W"
	initGrid(99, 99)
	rover := roverSetup()

	//When
	commandRover(rover, "bblbb")

	//Then
	actualX := rover.getX()

	if actualX != expectedX {
		t.Errorf("[X] Expected %d, actual %d", expectedX, actualX)
	}

	actualY := rover.getY()

	if actualY != expectedY {
		t.Errorf("[Y] Expected %d, actual %d", expectedY, actualY)
	}

	actualDirection := rover.getDirection()

	if actualDirection != expectedDirection {
		t.Errorf("[Direction] Expected %s, actual %s", expectedDirection, actualDirection)
	}
}
