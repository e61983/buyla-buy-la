package Buyla

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	Command_Open          string = "開團"
	Command_Close         string = "結單"
	Command_Show          string = "明細"
	Command_Help          string = "說明"
	Command_RD            string = "工程模式"
	TestCommand_Profile   string = "顯示使用者資訊"
	TestCommand_LIFF_Test string = "顯示LIFF測試選單"
	Surprise_1            string = "叫你們 RD 出來滴霸格"
)

type Bot struct {
	bot  *linebot.Client
	data *MetaData
}

func NewBot(channelSecret, channelToken, BaseUrl string, data *MetaData) (*Bot, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	return &Bot{bot: bot, data: data}, nil
}

func (this *Bot) replyMessage(replyToken string, messages ...linebot.SendingMessage) *linebot.ReplyMessageCall {
	messages[len(messages)-1].WithQuickReplies(linebot.NewQuickReplyItems(
		linebot.NewQuickReplyButton(
			"",
			linebot.NewMessageAction(Command_RD, "["+Command_RD+"]")),
	))
	return this.bot.ReplyMessage(replyToken, messages...)
}

func (this *Bot) replyText(replyToken, text string) error {
	if _, err := this.replyMessage(replyToken, linebot.NewTextMessage(text)).Do(); err != nil {
		return err
	}
	return nil
}

func getKeyWord(message string) string {
	r := regexp.MustCompile(`^(@([^[]*))?\[[\s\n\t ]*([^[]*)[\s\n\t ]*\][\s\n\t ]*(.*)`)

	message = strings.ReplaceAll(message, "［", "[")
	message = strings.ReplaceAll(message, "］", "]")

	token := r.FindAllStringSubmatch(message, 1)
	if len(token) == 0 {
		if strings.Contains(message, string(Surprise_1)) {
			return Surprise_1
		} else {
			return ""
		}
	}

	mentionName := strings.TrimSpace(token[0][2])
	keyword := strings.TrimSpace(token[0][3])
	others := strings.TrimSpace(token[0][4])

	log.Println("mentionName", mentionName)
	log.Println("keyword", keyword)
	log.Println("others", others)
	return keyword
}

func (this *Bot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	log.Println(message.Text)
	keyword := getKeyWord(message.Text)
	switch keyword {
	case TestCommand_Profile:
		if source.UserID != "" {
			profile, err := this.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return this.replyText(replyToken, err.Error())
			}
			if _, err := this.replyMessage(
				replyToken,
				linebot.NewTextMessage("使用者:"+profile.DisplayName),
			).Do(); err != nil {
				return err
			}
		} else {
			return this.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case TestCommand_LIFF_Test:
		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.ButtonComponent{
						Type: linebot.FlexComponentTypeButton,
						Action: &linebot.URIAction{
							Label: "目前訂單",
							URI:   "line://app/1602541695-bYKBPBe6",
						},
					},
				},
			},
		}
		if _, err := this.replyMessage(
			replyToken,
			linebot.NewFlexMessage("Menu-"+TestCommand_LIFF_Test, contents),
		).Do(); err != nil {
			return err
		}
	case Command_RD:

		newButtonComponent := func(displayName, command string) *linebot.ButtonComponent {
			return &linebot.ButtonComponent{
				Type:   linebot.FlexComponentTypeButton,
				Height: linebot.FlexButtonHeightTypeSm,
				Action: &linebot.MessageAction{
					Label: displayName,
					Text:  command,
				},
			}
		}

		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					newButtonComponent("DEBUG-"+TestCommand_Profile, "["+TestCommand_Profile+"]"),
					newButtonComponent("DEBUG-"+TestCommand_LIFF_Test, "["+TestCommand_LIFF_Test+"]"),
					newButtonComponent(Command_RD, "["+Command_RD+"]"),
					newButtonComponent(Command_Open, "["+Command_Open+"]"),
					newButtonComponent(Command_Close, "["+Command_Close+"]"),
					newButtonComponent(Command_Show, "["+Command_Show+"]"),
					newButtonComponent(Command_Help, "["+Command_Help+"]"),
					newButtonComponent(Surprise_1, Surprise_1),
				},
			},
		}

		if _, err := this.replyMessage(
			replyToken,
			linebot.NewFlexMessage("Menu-"+Command_RD, contents),
		).Do(); err != nil {
			return err
		}
	case Command_Open:
		return this.replyText(replyToken, "目前還不支援'["+keyword+"]'這個功能喔!")
	case Command_Close:
		return this.replyText(replyToken, "目前還不支援'["+keyword+"]'這個功能喔!")
	case Command_Help:
		functionFlex := 3
		descriptionFlex := 5

		newFunctionComponent := func(command, description string) *linebot.BoxComponent {
			return &linebot.BoxComponent{
				Type:    linebot.FlexComponentTypeBox,
				Layout:  linebot.FlexBoxLayoutTypeHorizontal,
				Spacing: linebot.FlexComponentSpacingTypeSm,
				Contents: []linebot.FlexComponent{
					&linebot.ButtonComponent{
						Type:    linebot.FlexComponentTypeButton,
						Style:   linebot.FlexButtonStyleTypePrimary,
						Flex:    &functionFlex,
						Height:  linebot.FlexButtonHeightTypeSm,
						Gravity: linebot.FlexComponentGravityTypeCenter,
						Action: &linebot.MessageAction{
							Label: command,
							Text:  "[" + command + "]",
						},
					},
					&linebot.TextComponent{
						Type:    linebot.FlexComponentTypeText,
						Size:    linebot.FlexTextSizeTypeSm,
						Align:   linebot.FlexComponentAlignTypeStart,
						Gravity: linebot.FlexComponentGravityTypeCenter,
						Flex:    &descriptionFlex,
						Text:    description,
					},
				},
			}
		}

		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeVertical,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type:   linebot.FlexComponentTypeText,
						Size:   linebot.FlexTextSizeTypeLg,
						Align:  linebot.FlexComponentAlignTypeCenter,
						Weight: linebot.FlexTextWeightTypeBold,
						Text:   "揪圑啦 的自我介紹",
					},
					&linebot.BoxComponent{
						Type:    linebot.FlexComponentTypeBox,
						Layout:  linebot.FlexBoxLayoutTypeVertical,
						Margin:  linebot.FlexComponentMarginTypeLg,
						Spacing: linebot.FlexComponentSpacingTypeSm,
						Contents: []linebot.FlexComponent{
							newFunctionComponent(Command_Open, "告訴大家有新的揪團!"),
							newFunctionComponent(Command_Close, "就是告訴大家下回請早的意思啦~"),
							newFunctionComponent(Command_Show, "看看大家訂了什麼"),
							newFunctionComponent(Command_Help, "跟大家再自我介紹一次"),
						},
					},
				},
			},
		}

		if _, err := this.replyMessage(
			replyToken,
			linebot.NewFlexMessage("Menu-"+Command_Help, contents),
		).Do(); err != nil {
			return err
		}
	case Command_Show:
		return this.replyText(replyToken, "目前還不支援'["+keyword+"]'這個功能喔!")
	case Surprise_1:
		if _, err := this.replyMessage(
			replyToken,
			linebot.NewStickerMessage("11537", "52002739"),
		).Do(); err != nil {
			return err
		}
	default:
		//log.Printf("Echo message to %s: %s", replyToken, message.Text)
		//if err := this.replyText(replyToken, message.Text); err != nil {
		//return err
		//}
	}
	return nil
}

func (this *Bot) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := this.bot.ParseRequest(r)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		log.Print(err)
		return
	}

	for _, event := range events {
		log.Printf("event: %v", event)
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := this.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Print(err)
				}
			case *linebot.ImageMessage:
			case *linebot.VideoMessage:
			case *linebot.AudioMessage:
			case *linebot.FileMessage:
			case *linebot.LocationMessage:
			case *linebot.StickerMessage:
			default:
			}
		case linebot.EventTypeFollow:
		case linebot.EventTypeUnfollow:
		case linebot.EventTypeJoin:
		case linebot.EventTypeLeave:
			log.Printf("Left: %v", event)
		case linebot.EventTypePostback:
		case linebot.EventTypeBeacon:
		default:
			log.Printf("Unknown event: %v", event)
		}
	}
}
