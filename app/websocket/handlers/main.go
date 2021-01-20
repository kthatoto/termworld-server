package handlers

import (
	"fmt"

	"github.com/kthatoto/termworld-server/app/models"
)

type Command struct {
	PlayerName string   `json:"playerName"`
	Command    string   `json:"command"`
	Options    []string `json:"options"`
	RequestId  string   `json:"requestId"`
}

type Response struct {
	RequestId string `json:"requestId"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
}

func Handle(currentUser *models.User, command Command) (Response, error) {
	resp := Response{
		RequestId: command.RequestId
		Success: false
		Message: ""
	}

	var playerModel models.PlayerModel
	player, err := playerModel.FindByName(*currentUser, command.PlayerName)
	if err != nil {
		resp.Message = fmt.Sprintf("player: %s is not found", command.PlayerName)
		return resp, err
	}

	switch command.Command {
	case "start":
		err = HandleStart(&player)
	default:
		resp.Message = fmt.Sprintf("command: %s is not found", command.Command)
		return resp, nil
	}

	if err != nil {
		resp.Message = fmt.Sprintf("error: %s", err)
		return resp, err
	}
	resp.Success = true
	return resp, nil
}

func HandleStart(player *models.Player) (error) {
	var playerModel models.PlayerModel
	err := playerModel.UpdateLive(player, true)
	if err != nil {
		return err
	}
	return nil
}
