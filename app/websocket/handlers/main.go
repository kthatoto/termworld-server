package handlers

import (
	"github.com/kthatoto/termworld-server/app/models"
)

type Command struct {
	PlayerName string   `json:"playerName"`
	Command    string   `json:"command"`
	Options    []string `json:"options"`
}

func Handle(currentUser *models.User, command Command) (error) {
	var playerModel models.PlayerModel
	player, err := playerModel.FindByName(*currentUser, command.PlayerName)
	if err != nil {
		return err
	}

	if (command.Command == "start") {
		err = HandleStart(&player)
	}

	if err != nil {
		return err
	}
	return nil
}

func HandleStart(player *models.Player) (error) {
	var playerModel models.PlayerModel
	err := playerModel.UpdateLive(player, true)
	if err != nil {
		return err
	}
	return nil
}
