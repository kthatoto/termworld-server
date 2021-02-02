package commands

import (
	"fmt"
	"errors"
	"io/ioutil"

	"github.com/kthatoto/termworld-server/app/models"
)

func Map(player *models.Player, resp *Response, options []string) error {
	if !player.Live {
		return errors.New(fmt.Sprintf("player[%s] は起動していません。まずstartコマンドで起動させてください", player.Name))
	}

	bytes, err := ioutil.ReadFile("./data/maps/map_0_0.json")
	if err != nil {
		return err
	}
	resp.Message = string(bytes)
	return nil
}
