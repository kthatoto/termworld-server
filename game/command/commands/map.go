package commands

import (
	"fmt"
	"errors"

	"github.com/kthatoto/termworld-server/app/models"
)

func Map(player *models.Player, resp *Response, options []string) error {
	if !player.Live {
		return errors.New(fmt.Sprintf("player[%s] は起動していません。まずstartコマンドで起動させてください", player.Name))
	}

	resp.Message = "map"
	return nil
}
