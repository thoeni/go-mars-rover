package mars_rover

import (
	"flag"
	"os"
	"testing"
)

var testGridXMax = 99
var testGridYMax = 99

func TestMain(m *testing.M) {
	flag.Parse()
	InitGrid(testGridXMax, testGridYMax)
	os.Exit(m.Run())
}

func TestRoverSetupShouldInitialiseTheRoverAt00DirectionN(t *testing.T) {
	t.Parallel()
	//Given
	expectedX := 0
	expectedY := 0
	expectedDirection := "N"

	//When
	rover := RoverSetup()

	//Then
	if rover == nil {
		t.Error("Rover has not been initialised")
	}

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

func TestMoveForwardFromOriginShouldIncreaseY(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()

	//Then
	if rover.GetY() != 1 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}
}

func TestTurnRightFromOriginShouldChangeTheDirectionToEast(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnRight()

	//Then
	if rover.GetDirection() != "E" {
		t.Error("Rover didn't face East when turning right from origin")
	}
}

func TestTurnRightTwiceFromOriginShouldChangeTheDirectionToSouth(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnRight()
	rover.TurnRight()

	//Then
	if rover.GetDirection() != "S" {
		t.Error("Rover didn't face South when turning right twice from origin")
	}
}

func TestTurnRightThreeTimesFromOriginShouldChangeTheDirectionToWest(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnRight()
	rover.TurnRight()
	rover.TurnRight()

	//Then
	if rover.GetDirection() != "W" {
		t.Error("Rover didn't face West when turning right three times from origin")
	}
}

func TestTurnRightFourTimesFromOriginShouldNotChangeTheDirection(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()
	initialDirection := rover.GetDirection()

	//When
	rover.TurnRight()
	rover.TurnRight()
	rover.TurnRight()
	rover.TurnRight()

	//Then
	if rover.GetDirection() != initialDirection {
		t.Error("Rover didn't change direction when turning right four times from origin")
	}
}

func TestMoveForwardTwiceFromOriginShouldIncreaseYTo2(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.MoveForward()

	//Then
	if rover.GetY() != 2 {
		t.Error("Rover didn't increase Y twice when moving forward twice from origin")
	}
}

func TestMoveForwardTurnRightAndMoveForwardShouldPositionRoverTo11(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.TurnRight()
	rover.MoveForward()

	//Then
	if rover.GetX() != 1 {
		t.Error("Rover didn't increase X when moving forward")
	}

	if rover.GetY() != 1 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}

	if rover.GetDirection() != "E" {
		t.Error("Rover isn't headed to East")
	}
}

func TestMoveForwardTwiceTurnRightAndMoveForwardTwiceShouldPositionRoverTo22(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.MoveForward()
	rover.TurnRight()
	rover.MoveForward()
	rover.MoveForward()

	//Then
	if rover.GetX() != 2 {
		t.Error("Rover didn't increase X when moving forward")
	}

	if rover.GetY() != 2 {
		t.Error("Rover didn't increase Y when moving forward from origin")
	}

	if rover.GetDirection() != "E" {
		t.Error("Rover isn't headed to East")
	}
}

func TestMoveBackwardsFromOriginShouldSetYTo99(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveBackwards()

	//Then
	if rover.GetY() != testGridYMax {
		t.Error("Rover didn't decrease Y when moving backwards from origin")
	}
}

func TestMoveForward100TimesShouldSetYTo0(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	for i := 0; i < testGridYMax+1; i++ {
		rover.MoveForward()
	}

	//Then
	if rover.GetY() != 0 {
		t.Error("Rover didn't go back to 0 when moving forward 100 times from origin")
	}
}

func TestTurnLeftFromOriginShouldChangeTheDirectionToWest(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnLeft()

	//Then
	if rover.GetDirection() != "W" {
		t.Error("Rover didn't face West when turning left from origin")
	}
}

func TestTurnLeftFromOriginAndMoveForwardShouldPositionTheRoverTox99y0(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnLeft()
	rover.MoveForward()

	//Then
	if rover.GetDirection() != "W" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.GetX() != 99 {
		t.Error("Rover didn't go around the board")
	}
}

func TestTurnRightFromOriginAndMoveBackwardsTwiceShouldPositionTheRoverTox98y0(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.TurnRight()
	rover.MoveBackwards()
	rover.MoveBackwards()

	//Then
	if rover.GetDirection() != "E" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.GetX() != testGridXMax-1 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardAndBackwardShouldKeepTheRoverInTheSamePosition(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.MoveBackwards()

	//Then
	if rover.GetDirection() != "N" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.GetY() != 0 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardRotate180AndMoveForwardShouldKeepTheRoverInTheSamePosition(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.TurnLeft()
	rover.TurnLeft()
	rover.MoveForward()

	//Then
	if rover.GetDirection() != "S" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.GetY() != 0 {
		t.Error("Rover didn't go around the board")
	}
}

func TestMoveForwardRotate180AndMoveBackwardsShouldMoveTheRoverTo(t *testing.T) {
	t.Parallel()
	//Given
	rover := RoverSetup()

	//When
	rover.MoveForward()
	rover.TurnLeft()
	rover.TurnLeft()
	rover.MoveBackwards()

	//Then
	if rover.GetDirection() != "S" {
		t.Error("Rover didn't face West when turning left from origin")
	}

	if rover.GetY() != 2 {
		t.Error("Rover didn't go around the board")
	}
}
