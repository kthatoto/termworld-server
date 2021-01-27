package commands

import (
	"fmt"
	"errors"

	"github.com/kthatoto/termworld-server/app/models"
)

func Move(player *models.Player, resp *Response, options []string) error {
	direction := options[0]
	fmt.Printf("%s move command recieved!", direction)
	if !(direction == "up" || direction == "down" || direction == "left" || direction == "right") {
		return errors.New(fmt.Sprintf("[%s] is invalid direction (The direction should be [up, down, left or right])", direction))
	}
	return nil
}
