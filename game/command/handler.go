package command

import (
	"fmt"

	"github.com/kthatoto/termworld-server/app/models"
	"github.com/kthatoto/termworld-server/game/command/commands"
)

type Command struct {
	PlayerName string   `json:"playerName"`
	Command    string   `json:"command"`
	Options    []string `json:"options"`
	RequestId  string   `json:"requestId"`
}

func Handle(currentUser *models.User, command Command) (commands.Response, error) {
	resp := commands.Response{
		RequestId: command.RequestId,
		Success: false,
		Message: "",
	}

	var playerModel models.PlayerModel
	player, err := playerModel.FindByName(*currentUser, command.PlayerName)
	if err != nil {
		resp.Message = fmt.Sprintf("player: %s is not found", command.PlayerName)
		return resp, err
	}

	switch command.Command {
	case "start":
		err = commands.Start(&player, &resp, command.Options)
	case "stop":
		err = commands.Stop(&player, &resp, command.Options)
	case "move":
		err = commands.Move(&player, &resp, command.Options)
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
