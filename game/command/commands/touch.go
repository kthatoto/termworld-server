package commands

import (
	"fmt"

	"github.com/kthatoto/termworld-server/app/models"
)

func Touch(player *models.Player, resp *Response, options []string) error {
	if (player.Live) {
		resp.Message = "live"
	} else {
		resp.Message = fmt.Sprintf("player[%s] は起動していません。まずstartコマンドで起動させてください", player.Name)
	}
	return nil
}
