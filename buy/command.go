package buy

import (
	"errors"
	"log"
	"strings"
)

type CommandType string

const (
	CommandTypeOpenNewBuyLa CommandType = "開團"
	CommandTypeWant         CommandType = "我要"
	CommandTypeCloseBuyLa   CommandType = "結單"
	CommandTypeShowRecord   CommandType = "明細"
	CommandTypeMeToo_v1     CommandType = "咪兔"
	CommandTypeMeToo_v2     CommandType = "+1"
	CommandTypeHelp         CommandType = "說明"
	CommandTypeCancel       CommandType = "我不要了"
	CommandTypeAttach       CommandType = "加訂"
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
type CancelCommand struct{ UserID string }
type AttchCommand struct{ UserID, Goods string }
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

func NewCancelCommand(userID string) Command {
	return &CancelCommand{UserID: userID}
}

func NewAttchCommand(userID, goods string) Command {
	return &AttchCommand{UserID: userID, Goods: goods}
}

func NewRDDebugCommand(userID string) Command {
	return &RDDebugCommand{UserID: userID}
}

func ParseCommand(userID, message string) (command Command, err error) {
	err = nil
	command = nil

	token := parse.FindAllStringSubmatch(message, 1)
	if len(token) == 0 {
		if strings.Contains(message, string(CommandTypeRDDebug)) {
			return NewRDDebugCommand(userID), nil
		} else {
			err = errors.New("Not a Command")
			return
		}
	}

	mentionName := strings.TrimSpace(token[0][2])
	keyword := CommandType(strings.TrimSpace(token[0][3]))
	others := strings.TrimSpace(token[0][4])

	log.Println("mentionName", mentionName)
	log.Println("keyword", keyword)
	log.Println("others", others)

	switch keyword {
	case CommandTypeOpenNewBuyLa:
		command = NewOpenNewBuyLaCommand(userID, others)
	case CommandTypeWant:
		command = NewWantCommand(userID, others)
	case CommandTypeCloseBuyLa:
		command = NewCloseBuyLaCommand(userID)
	case CommandTypeShowRecord:
		command = NewShowRecordCommand(userID)
	case CommandTypeMeToo_v1:
		fallthrough
	case CommandTypeMeToo_v2:
		command = NewMeTooCommand(userID, mentionName)
	case CommandTypeHelp:
		command = NewHelpCommand(userID)
	case CommandTypeAttach:
		command = NewAttchCommand(userID, others)
	case CommandTypeCancel:
		command = NewCancelCommand(userID)
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
func (c *CancelCommand) Command()       {}
func (c *RDDebugCommand) Command()      {}
func (c *AttchCommand) Command()        {}
