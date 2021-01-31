package commands

import (
	"fmt"

	"github.com/kthatoto/termworld-server/app/models"
)

func Start(player *models.Player, resp *Response, options []string) error {
	var playerModel models.PlayerModel
	err := playerModel.UpdateLive(player, true)
	if err != nil {
		return err
	}
	err = playerModel.StartPlayer(player)
	if err != nil {
		return err
	}

	resp.Message = fmt.Sprintf("player[%s] started", player.Name)
	return nil
}

