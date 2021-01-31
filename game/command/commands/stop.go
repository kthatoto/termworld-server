package commands

import (
	"fmt"

	"github.com/kthatoto/termworld-server/app/models"
)

func Stop(player *models.Player, resp *Response, options []string) error {
	var playerModel models.PlayerModel
	err := playerModel.UpdateLive(player, false)
	if err != nil {
		return err
	}

	resp.Message = fmt.Sprintf("player[%s] stopped", player.Name)
	return nil
}
