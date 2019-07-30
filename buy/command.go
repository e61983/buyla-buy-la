package buy

import (
	"errors"
	"strings"
)

type CommandType string

const (
	CommandTypeOpenNewBuyLa CommandType = "開團"
	CommandTypeWant         CommandType = "我要"
	CommandTypeCloseBuyLa   CommandType = "結單"
	CommandTypeShowRecord   CommandType = "明細"
	CommandTypeMeToo        CommandType = "咪兔"
	CommandTypeHelp         CommandType = "說明"
	CommandTypeRDDebug      CommandType = "叫你們 RD 出來滴霸格!!!"
)

type Command interface {
	Command()
}

type OpenNewBuyLaCommand struct{ UserID, ShopName string }
type WantCommand struct{ UserID, Goods string }
type CloseBuyLaCommand struct{ UserID string }
type ShowRecordCommand struct{ UserID string }
type MeTooCommand struct{ UserID, TargetName string }
type HelpCommand struct{ UserID string }
type RDDebugCommand struct{ UserID string }

func NewOpenNewBuyLaCommand(userID, shop string) Command {
	return &OpenNewBuyLaCommand{UserID: userID, ShopName: shop}
}

func NewWantCommand(userID, goods string) Command {
	return &WantCommand{UserID: userID, Goods: goods}
}

func NewCloseBuyLaCommand(userID string) Command {
	return &CloseBuyLaCommand{UserID: userID}
}

func NewShowRecordCommand(userID string) Command {
	return &ShowRecordCommand{UserID: userID}
}

func NewMeTooCommand(userID, targetName string) Command {
	return &MeTooCommand{UserID: userID, TargetName: targetName}
}

func NewHelpCommand(userID string) Command {
	return &HelpCommand{UserID: userID}
}

func NewRDDebugCommand(userID string) Command {
	return &RDDebugCommand{UserID: userID}
}

func ParseCommand(userID, message string) (command Command, err error) {
	err = nil
	command = nil

	token := parse.FindAllStringSubmatch(message, 1)
	if len(token) == 0 {
		err = errors.New("Not a Command")
	}

	mentionName := strings.TrimSpace(token[0][2])
	keyword := CommandType(strings.TrimSpace(token[0][3]))
	others := strings.TrimSpace(token[0][4])

	if userID == "" {
		err = errors.New("User ID is Empty")
	}

	if err != nil {
		return
	}

	switch keyword {
	case CommandTypeOpenNewBuyLa:
		if others == "" {
			err = errors.New("Shop Name is Empty")
		}
		command = NewOpenNewBuyLaCommand(userID, others)
	case CommandTypeWant:
		if others == "" {
			err = errors.New("Goods is Empty")
		}
		command = NewWantCommand(userID, others)
	case CommandTypeCloseBuyLa:
		command = NewCloseBuyLaCommand(userID)
	case CommandTypeShowRecord:
		command = NewShowRecordCommand(userID)
	case CommandTypeMeToo:
		if mentionName == "" {
			err = errors.New("MentionName is Empty")
		}
		command = NewMeTooCommand(userID, mentionName)
	case CommandTypeHelp:
		command = NewHelpCommand(userID)
	case CommandTypeRDDebug:
		command = NewRDDebugCommand(userID)
	default:
	}

	return
}

func (c *OpenNewBuyLaCommand) Command() {}
func (c *WantCommand) Command()         {}
func (c *CloseBuyLaCommand) Command()   {}
func (c *ShowRecordCommand) Command()   {}
func (c *MeTooCommand) Command()        {}
func (c *HelpCommand) Command()         {}
func (c *RDDebugCommand) Command()      {}
