package mars_rover

import "testing"

func TestCommandRoverFFRFF(t *testing.T) {
	t.Parallel()
	//Given
	expectedX := 2
	expectedY := 2
	expectedDirection := "E"
	InitGrid(99, 99)
	rover := RoverSetup()

	//When
	CommandRover(rover, "ffrff")

	//Then
	actualX := rover.GetX()

	if actualX != expectedX {
		t.Errorf("[X] Expected %d, actual %d", expectedX, actualX)
	}

	actualY := rover.GetY()

	if actualY != expectedY {
		t.Errorf("[Y] Expected %d, actual %d", expectedY, actualY)
	}

	actualDirection := rover.GetDirection()

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
	InitGrid(99, 99)
	rover := RoverSetup()

	//When
	CommandRover(rover, "bblbb")

	//Then
	actualX := rover.GetX()

	if actualX != expectedX {
		t.Errorf("[X] Expected %d, actual %d", expectedX, actualX)
	}

	actualY := rover.GetY()

	if actualY != expectedY {
		t.Errorf("[Y] Expected %d, actual %d", expectedY, actualY)
	}

	actualDirection := rover.GetDirection()

	if actualDirection != expectedDirection {
		t.Errorf("[Direction] Expected %s, actual %s", expectedDirection, actualDirection)
	}
}
