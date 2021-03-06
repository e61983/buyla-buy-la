package Buyla

import (
	"bytes"
	"encoding/json"
	"github.com/line/line-bot-sdk-go/linebot"
	"log"
	"net/http"
	"regexp"
	"strings"
)

const (
	Command_Open                 string = "開團"
	Command_Close                string = "結單"
	Command_Show                 string = "明細"
	Command_Help                 string = "說明"
	Command_RD                   string = "工程模式"
	TestCommand_Profile          string = "顯示使用者資訊"
	TestCommand_AddTestRecord    string = "建立測試點單"
	TestCommand_DeleteTestRecord string = "刪除測試點單"
	TestCommand_LIFF_Test        string = "顯示LIFF測試選單"
	Surprise_1                   string = "叫你們 RD 出來滴霸格"
)

type Bot struct {
	bot     *linebot.Client
	data    *MetaData
	baseUrl string
}

func getUID(source *linebot.EventSource) string {
	return source.UserID
}

func getGID(source *linebot.EventSource) string {
	if source.GroupID == "" {
		// Just For test
		//return "test"
		return ""
	} else {
		return source.GroupID
	}
}

func getRecordContents(group *Group) linebot.FlexContainer {
	newItemComponent := func(record *Record) *linebot.BoxComponent {
		box := &linebot.BoxComponent{
			Type:    linebot.FlexComponentTypeBox,
			Layout:  linebot.FlexBoxLayoutTypeVertical,
			Spacing: linebot.FlexComponentSpacingTypeSm,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Size:    linebot.FlexTextSizeTypeSm,
					Weight:  linebot.FlexTextWeightTypeBold,
					Align:   linebot.FlexComponentAlignTypeStart,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					//Flex:    &descriptionFlex,
					Text: record.UserName,
				},
			},
		}

		for _, v := range record.Order.List {
			log.Println("[SHOW]", record.UserName, v)
			box.Contents = append(box.Contents,
				&linebot.TextComponent{
					Type:    linebot.FlexComponentTypeText,
					Size:    linebot.FlexTextSizeTypeSm,
					Align:   linebot.FlexComponentAlignTypeStart,
					Gravity: linebot.FlexComponentGravityTypeCenter,
					//Flex:    &descriptionFlex,
					Text: "•" + v.ItemName + " " + v.SweetnessLevel + " " + v.AmountOfIce + " x " + v.Number,
				})
		}
		return box
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
					Text:   "明細",
				},
			},
		},
	}

	if len(group.Records) > 0 {
		for _, record := range group.Records {
			contents.Body.Contents = append(contents.Body.Contents,
				&linebot.SeparatorComponent{
					Type:   linebot.FlexComponentTypeSeparator,
					Margin: linebot.FlexComponentMarginTypeXxl,
				},
				newItemComponent(record))
		}
	} else {
		contents.Body.Contents = append(contents.Body.Contents,
			&linebot.SeparatorComponent{
				Type:   linebot.FlexComponentTypeSeparator,
				Margin: linebot.FlexComponentMarginTypeXxl,
			},
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Size:  linebot.FlexTextSizeTypeSm,
				Align: linebot.FlexComponentAlignTypeStart,
				Text:  "好像什麼也沒有...",
			})
	}
	return contents
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
	log.Println("[COMMAND]", keyword, mentionName, others)
	return keyword
}

func NewBot(channelSecret, channelToken, BaseUrl string, data *MetaData) (*Bot, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	return &Bot{bot: bot, baseUrl: BaseUrl, data: data}, nil
}

func (this *Bot) getProfile(gid, uid string) (*linebot.UserProfileResponse, error) {
	if gid != "" && gid != "test" {
		return this.bot.GetGroupMemberProfile(gid, uid).Do()
	} else {
		return this.bot.GetProfile(uid).Do()
	}
}

func (this *Bot) replyMessage(replyToken string, messages ...linebot.SendingMessage) *linebot.ReplyMessageCall {
	return this.bot.ReplyMessage(replyToken, messages...)
}

func (this *Bot) replyText(replyToken, text string) error {
	if _, err := this.replyMessage(replyToken, linebot.NewTextMessage(text)).Do(); err != nil {
		return err
	}
	return nil
}

func (this *Bot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	keyword := getKeyWord(message.Text)
	switch keyword {
	case TestCommand_Profile:
		uid := getUID(source)
		gid := getGID(source)
		profile, err := this.getProfile(gid, uid)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		if _, err := this.replyMessage(
			replyToken,
			linebot.NewTextMessage("使用者:"+profile.DisplayName),
		).Do(); err != nil {
			return err
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
	case TestCommand_AddTestRecord:
		uid := getUID(source)
		gid := getGID(source)
		profile, err := this.getProfile(gid, uid)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		displayName := profile.DisplayName
		testRecord := &Record{
			UserName: displayName,
			Order: &OrderItems{
				List: []*OrderItem{
					&OrderItem{"休假", "跟你一樣甜", "不要冰的", "很多很多"},
				},
			},
		}
		testRecordJSON := new(bytes.Buffer)
		json.NewEncoder(testRecordJSON).Encode(&testRecord)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		client := http.Client{}
		//TODO show usgin variable
		url := this.baseUrl + "/api/v1/" + gid + "/order/" + uid
		req, err := http.NewRequest(http.MethodPost, url, testRecordJSON)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		_, err = client.Do(req)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		if _, err = this.replyMessage(
			replyToken,
			linebot.NewTextMessage(displayName+"建立了測試訂點單"),
		).Do(); err != nil {
			return err
		}
	case TestCommand_DeleteTestRecord:
		uid := getUID(source)
		gid := getGID(source)
		client := http.Client{}
		url := this.baseUrl + "/api/v1/" + gid + "/order/" + uid
		req, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		res, err := client.Do(req)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		} else {
			log.Println(res)
		}

		profile, err := this.getProfile(gid, uid)
		if err != nil {
			return this.replyText(replyToken, err.Error())
		}
		if _, err = this.replyMessage(
			replyToken,
			linebot.NewTextMessage("刪除了"+profile.DisplayName+"的測試訂點單"),
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
					newButtonComponent("TEST-"+TestCommand_Profile, "["+TestCommand_Profile+"]"),
					newButtonComponent("TEST-"+TestCommand_LIFF_Test, "["+TestCommand_LIFF_Test+"]"),
					newButtonComponent("TEST-"+TestCommand_AddTestRecord, "["+TestCommand_AddTestRecord+"]"),
					newButtonComponent("TEST-"+TestCommand_DeleteTestRecord, "["+TestCommand_DeleteTestRecord+"]"),
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
		gid := getGID(source)
		if gid == "" {
			return this.replyText(replyToken, "'["+keyword+"]'只能在群組裡面使用喔!")
		}

		if _, ok := this.data.Groups[gid]; !ok {
			this.data.Groups[gid] = NewGroup()
			log.Println("[CREATE]", gid)
		} else {
			if this.data.Groups[gid].IsOpen == false {
				this.data.Groups[gid] = NewGroup()
				log.Println("[CLEAN]", gid)
			} else {
				log.Println("[ABORT]", gid)
				return this.replyText(replyToken, "已經在開了喔~!")
			}
		}

		this.data.Groups[gid].IsOpen = true
		log.Println("[OPEN]", gid)
		return this.replyText(replyToken, "開團啦~~!!!!! ")

	case Command_Close:
		gid := getGID(source)
		if gid == "" {
			return this.replyText(replyToken, "'["+keyword+"]'只能在群組裡面使用喔!")
		}
		if _, ok := this.data.Groups[gid]; ok && this.data.Groups[gid].IsOpen == true {
			this.data.Groups[gid].IsOpen = false
			log.Println("[CLOSE]", gid)
		} else {
			return this.replyText(replyToken, "現在還沒有開始揪團~\n 大家都在等你開喔~!! XD")
		}
		if _, err := this.replyMessage(
			replyToken,
			linebot.NewTextMessage("熱騰騰的明細出來啦~~"),
			linebot.NewFlexMessage("明細", getRecordContents(this.data.Groups[gid])),
		).Do(); err != nil {
			return err
		}
		//return this.replyText(replyToken, "結單啦~!")
	case Command_Show:
		gid := getGID(source)
		log.Println("[SHOW]", gid)
		if gid == "" {
			return this.replyText(replyToken, "'["+keyword+"]'只能在群組裡面使用喔!")
		}

		var group *Group
		if _, ok := this.data.Groups[gid]; ok {
			group = this.data.Groups[gid]
		} else {
			return this.replyText(replyToken, "還沒有開團過喔~!\n")
		}

		if _, err := this.replyMessage(
			replyToken,
			linebot.NewFlexMessage("明細", getRecordContents(group)),
		).Do(); err != nil {
			return err
		}
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
