package commands

import (
	"fmt"
	"errors"

	"github.com/kthatoto/termworld-server/app/models"
)

func Move(player *models.Player, resp *Response, options []string) error {
	if !player.Live {
		return errors.New(fmt.Sprintf("player[%s] は起動していません。まずstartコマンドで起動させてください", player.Name))
	}
	if len(options) == 0 {
		return errors.New("moveする方向が必要です (up, down, left, right)")
	}

	direction := options[0]
	if !(direction == "up" || direction == "down" || direction == "left" || direction == "right") {
		return errors.New(fmt.Sprintf("[%s] は不正な方向です。up, down, left, rightから入力してください", direction))
	}
	dx, dy := 0, 0
	switch direction {
	case "up":
		dy = 1
	case "down":
		dy = -1
	case "left":
		dx = -1
	case "right":
		dx = 1
	}

	var playerModel models.PlayerModel
	playerModel.Move(player, dx, dy)
	resp.Message = fmt.Sprintf("player[%s] moved", player.Name)
	return nil
}
