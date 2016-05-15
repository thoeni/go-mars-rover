package marsrover

import (
	"flag"
	"os"
	"testing"
)

var testGridXMax = 99
var testGridYMax = 99

func TestMain(m *testing.M) {
	flag.Parse()
	initGrid(testGridXMax, testGridYMax)
	os.Exit(m.Run())
}

func TestroverSetupShouldInitialiseTheRoverAt00DirectionN(t *testing.T) {
	t.Parallel()
	//Given
	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	//When
	rover := roverSetup()

	//Then
	if rover == nil {
		t.Error("Rover has not been initialised")
	}

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

func TestMoveForwardFromOriginShouldIncreaseY(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()

	//Then
	if rover.getY() != 1 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}
}

func TestTurnRightFromOriginShouldChangeTheDirectionToEast(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnRight()

	//Then
	if rover.getDirection() != "E" {
		t.Error("Rover didn't face East when turning right from origin")
	}
}

func TestTurnRightTwiceFromOriginShouldChangeTheDirectionToSouth(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnRight()
	rover.turnRight()

	//Then
	if rover.getDirection() != "S" {
		t.Error("Rover didn't face South when turning right twice from origin")
	}
}

func TestTurnRightThreeTimesFromOriginShouldChangeTheDirectionToWest(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnRight()
	rover.turnRight()
	rover.turnRight()

	//Then
	if rover.getDirection() != "W" {
		t.Error("Rover didn't face West when turning right three times from origin")
	}
}

func TestTurnRightFourTimesFromOriginShouldNotChangeTheDirection(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()
	initialDirection := rover.getDirection()

	//When
	rover.turnRight()
	rover.turnRight()
	rover.turnRight()
	rover.turnRight()

	//Then
	if rover.getDirection() != initialDirection {
		t.Error("Rover didn't change direction when turning right four times from origin")
	}
}

func TestMoveForwardTwiceFromOriginShouldIncreaseYTo2(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.moveForward()

	//Then
	if rover.getY() != 2 {
		t.Error("Rover didn't increase Y twice when moving forward twice from origin")
	}
}

func TestMoveForwardTurnRightAndMoveForwardShouldPositionRoverTo11(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.turnRight()
	rover.moveForward()

	//Then
	if rover.getX() != 1 {
		t.Error("Rover didn't increase X when moving forward")
	}

	if rover.getY() != 1 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}

	if rover.getDirection() != "E" {
		t.Error("Rover isn't headed to East")
	}
}

func TestMoveForwardTwiceTurnRightAndMoveForwardTwiceShouldPositionRoverTo22(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.moveForward()
	rover.turnRight()
	rover.moveForward()
	rover.moveForward()

	//Then
	if rover.getX() != 2 {
		t.Error("Rover didn't increase X when moving forward")
	}

	if rover.getY() != 2 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}

	if rover.getDirection() != "E" {
		t.Error("Rover isn't headed to East")
	}
}

func TestMoveBackwardsFromOriginShouldSetYTo99(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveBackwards()

	//Then
	if rover.getY() != testGridYMax {
		t.Error("Rover didn't decrease Y when moving backwards from origin")
	}
}

func TestMoveForward100TimesShouldSetYTo0(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	for i := 0; i < testGridYMax+1; i++ {
		rover.moveForward()
	}

	//Then
	if rover.getY() != 0 {
		t.Error("Rover didn't go back to 0 when moving forward 100 times from origin")
	}
}

func TestTurnLeftFromOriginShouldChangeTheDirectionToWest(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnLeft()

	//Then
	if rover.getDirection() != "W" {
		t.Error("Rover didn't face West when turning left from origin")
	}
}

func TestTurnLeftFromOriginAndMoveForwardShouldPositionTheRoverTox99y0(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnLeft()
	rover.moveForward()

	//Then
	if rover.getDirection() != "W" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.getX() != 99 {
		t.Error("Rover didn't go around the board")
	}
}

func TestTurnRightFromOriginAndMoveBackwardsTwiceShouldPositionTheRoverTox98y0(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.turnRight()
	rover.moveBackwards()
	rover.moveBackwards()

	//Then
	if rover.getDirection() != "E" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.getX() != testGridXMax-1 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardAndBackwardShouldKeepTheRoverInTheSamePosition(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.moveBackwards()

	//Then
	if rover.getDirection() != "N" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.getY() != 0 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardRotate180AndMoveForwardShouldKeepTheRoverInTheSamePosition(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.turnLeft()
	rover.turnLeft()
	rover.moveForward()

	//Then
	if rover.getDirection() != "S" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.getY() != 0 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardRotate180AndMoveBackwardsShouldMoveTheRoverTo(t *testing.T) {
	t.Parallel()
	//Given
	rover := roverSetup()

	//When
	rover.moveForward()
	rover.turnLeft()
	rover.turnLeft()
	rover.moveBackwards()

	//Then
	if rover.getDirection() != "S" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.getY() != 2 {
		t.Error("Rover didn't go around the board")
	}
}
