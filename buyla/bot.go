package Buyla

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
)

type Bot struct {
	bot         *linebot.Client
	BaseUrl     string
	DownloadDir string
}

func NewBot(channelSecret, channelToken, BaseUrl string) (*Bot, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	DownloadDir := filepath.Join(filepath.Dir(os.Args[0]), "dl")
	_, err = os.Stat(DownloadDir)
	if err != nil {
		if err := os.Mkdir(DownloadDir, 0777); err != nil {
			return nil, err
		}
	}
	log.Println("create", DownloadDir)

	return &Bot{bot: bot, BaseUrl: BaseUrl, DownloadDir: DownloadDir}, nil
}

func (this *Bot) replyText(replyToken, text string) error {
	if _, err := this.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(text),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (this *Bot) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	switch message.Text {
	case "profile":
		if source.UserID != "" {
			profile, err := this.bot.GetProfile(source.UserID).Do()
			if err != nil {
				return this.replyText(replyToken, err.Error())
			}
			if _, err := this.bot.ReplyMessage(
				replyToken,
				linebot.NewTextMessage("Display name: "+profile.DisplayName),
				linebot.NewTextMessage("Status message: "+profile.StatusMessage),
			).Do(); err != nil {
				return err
			}
		} else {
			return this.replyText(replyToken, "Bot can't use profile API without user ID")
		}
	case "buttons":
		imageURL := this.BaseUrl + "/static/buttons/1040.jpg"
		template := linebot.NewButtonsTemplate(
			imageURL, "My button sample", "Hello, my button",
			linebot.NewURIAction("Go to line.me", "https://line.me"),
			linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", "hello こんにちは"),
			linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
			linebot.NewMessageAction("Say message", "Rice=米"),
		)
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Buttons alt text", template),
		).Do(); err != nil {
			return err
		}
	case "confirm":
		template := linebot.NewConfirmTemplate(
			"Do it?",
			linebot.NewMessageAction("Yes", "Yes!"),
			linebot.NewMessageAction("No", "No!"),
		)
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Confirm alt text", template),
		).Do(); err != nil {
			return err
		}
	case "carousel":
		imageURL := this.BaseUrl + "/static/buttons/1040.jpg"
		template := linebot.NewCarouselTemplate(
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewURIAction("Go to line.me", "https://line.me"),
				linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewCarouselColumn(
				imageURL, "hoge", "fuga",
				linebot.NewPostbackAction("言 hello2", "hello こんにちは", "hello こんにちは", ""),
				linebot.NewMessageAction("Say message", "Rice=米"),
			),
		)
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "image carousel":
		imageURL := this.BaseUrl + "/static/buttons/1040.jpg"
		template := linebot.NewImageCarouselTemplate(
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewURIAction("Go to LINE", "https://line.me"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewPostbackAction("Say hello1", "hello こんにちは", "", ""),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewMessageAction("Say message", "Rice=米"),
			),
			linebot.NewImageCarouselColumn(
				imageURL,
				linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
			),
		)
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Image carousel alt text", template),
		).Do(); err != nil {
			return err
		}
	case "datetime":
		template := linebot.NewButtonsTemplate(
			"", "", "Select date / time !",
			linebot.NewDatetimePickerAction("date", "DATE", "date", "", "", ""),
			linebot.NewDatetimePickerAction("time", "TIME", "time", "", "", ""),
			linebot.NewDatetimePickerAction("datetime", "DATETIME", "datetime", "", "", ""),
		)
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTemplateMessage("Datetime pickers alt text", template),
		).Do(); err != nil {
			return err
		}
	case "flex":
		// {
		//   "type": "bubble",
		//   "body": {
		//     "type": "box",
		//     "layout": "horizontal",
		//     "contents": [
		//       {
		//         "type": "text",
		//         "text": "Hello,"
		//       },
		//       {
		//         "type": "text",
		//         "text": "World!"
		//       }
		//     ]
		//   }
		// }
		contents := &linebot.BubbleContainer{
			Type: linebot.FlexContainerTypeBubble,
			Body: &linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "Hello,",
					},
					&linebot.TextComponent{
						Type: linebot.FlexComponentTypeText,
						Text: "World!",
					},
				},
			},
		}
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "flex carousel":
		// {
		//   "type": "carousel",
		//   "contents": [
		//     {
		//       "type": "bubble",
		//       "body": {
		//         "type": "box",
		//         "layout": "vertical",
		//         "contents": [
		//           {
		//             "type": "text",
		//             "text": "First bubble"
		//           }
		//         ]
		//       }
		//     },
		//     {
		//       "type": "bubble",
		//       "body": {
		//         "type": "box",
		//         "layout": "vertical",
		//         "contents": [
		//           {
		//             "type": "text",
		//             "text": "Second bubble"
		//           }
		//         ]
		//       }
		//     }
		//   ]
		// }
		contents := &linebot.CarouselContainer{
			Type: linebot.FlexContainerTypeCarousel,
			Contents: []*linebot.BubbleContainer{
				{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "First bubble",
							},
						},
					},
				},
				{
					Type: linebot.FlexContainerTypeBubble,
					Body: &linebot.BoxComponent{
						Type:   linebot.FlexComponentTypeBox,
						Layout: linebot.FlexBoxLayoutTypeVertical,
						Contents: []linebot.FlexComponent{
							&linebot.TextComponent{
								Type: linebot.FlexComponentTypeText,
								Text: "Second bubble",
							},
						},
					},
				},
			},
		}
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "flex json":
		jsonString := `{
  "type": "bubble",
  "hero": {
    "type": "image",
    "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/01_1_cafe.png",
    "size": "full",
    "aspectRatio": "20:13",
    "aspectMode": "cover",
    "action": {
      "type": "uri",
      "uri": "http://linecorp.com/"
    }
  },
  "body": {
    "type": "box",
    "layout": "vertical",
    "contents": [
      {
        "type": "text",
        "text": "Brown Cafe",
        "weight": "bold",
        "size": "xl"
      },
      {
        "type": "box",
        "layout": "baseline",
        "margin": "md",
        "contents": [
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/review_gold_star_28.png"
          },
          {
            "type": "icon",
            "size": "sm",
            "url": "https://scdn.line-thiss.com/n/channel_devcenter/img/fx/review_gray_star_28.png"
          },
          {
            "type": "text",
            "text": "4.0",
            "size": "sm",
            "color": "#999999",
            "margin": "md",
            "flex": 0
          }
        ]
      },
      {
        "type": "box",
        "layout": "vertical",
        "margin": "lg",
        "spacing": "sm",
        "contents": [
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Place",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "Miraina Tower, 4-1-6 Shinjuku, Tokyo",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          },
          {
            "type": "box",
            "layout": "baseline",
            "spacing": "sm",
            "contents": [
              {
                "type": "text",
                "text": "Time",
                "color": "#aaaaaa",
                "size": "sm",
                "flex": 1
              },
              {
                "type": "text",
                "text": "10:00 - 23:00",
                "wrap": true,
                "color": "#666666",
                "size": "sm",
                "flex": 5
              }
            ]
          }
        ]
      }
    ]
  },
  "footer": {
    "type": "box",
    "layout": "vertical",
    "spacing": "sm",
    "contents": [
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "CALL",
          "uri": "https://linecorp.com"
        }
      },
      {
        "type": "button",
        "style": "link",
        "height": "sm",
        "action": {
          "type": "uri",
          "label": "WEBSITE",
          "uri": "https://linecorp.com",
          "altUri": {
            "desktop": "https://line.me/ja/download"
          }
        }
      },
      {
        "type": "spacer",
        "size": "sm"
      }
    ],
    "flex": 0
  }
}`
		contents, err := linebot.UnmarshalFlexMessageJSON([]byte(jsonString))
		if err != nil {
			return err
		}
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewFlexMessage("Flex message alt text", contents),
		).Do(); err != nil {
			return err
		}
	case "imagemap":
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewImagemapMessage(
				this.BaseUrl+"/static/rich",
				"Imagemap alt text",
				linebot.ImagemapBaseSize{Width: 1040, Height: 1040},
				linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520}),
			),
		).Do(); err != nil {
			return err
		}
	case "imagemap video":
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewImagemapMessage(
				this.BaseUrl+"/static/rich",
				"Imagemap with video alt text",
				linebot.ImagemapBaseSize{Width: 1040, Height: 1040},
				linebot.NewMessageImagemapAction("URANAI!", linebot.ImagemapArea{X: 520, Y: 520, Width: 520, Height: 520}),
			).WithVideo(&linebot.ImagemapVideo{
				OriginalContentURL: this.BaseUrl + "/static/imagemap/video.mp4",
				PreviewImageURL:    this.BaseUrl + "/static/imagemap/preview.jpg",
				Area:               linebot.ImagemapArea{X: 280, Y: 385, Width: 480, Height: 270},
				ExternalLink:       &linebot.ImagemapVideoExternalLink{LinkURI: "https://line.me", Label: "LINE"},
			}),
		).Do(); err != nil {
			return err
		}
	case "quick":
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage("Select your favorite food category or send me your location!").
				WithQuickReplies(linebot.NewQuickReplyItems(
					linebot.NewQuickReplyButton(
						this.BaseUrl+"/static/quick/sushi.png",
						linebot.NewMessageAction("Sushi", "Sushi")),
					linebot.NewQuickReplyButton(
						this.BaseUrl+"/static/quick/tempura.png",
						linebot.NewMessageAction("Tempura", "Tempura")),
					linebot.NewQuickReplyButton(
						"",
						linebot.NewLocationAction("Send location")),
				)),
		).Do(); err != nil {
			return err
		}
	case "bye":
		switch source.Type {
		case linebot.EventSourceTypeUser:
			return this.replyText(replyToken, "Bot can't leave from 1:1 chat")
		case linebot.EventSourceTypeGroup:
			if err := this.replyText(replyToken, "Leaving group"); err != nil {
				return err
			}
			if _, err := this.bot.LeaveGroup(source.GroupID).Do(); err != nil {
				return this.replyText(replyToken, err.Error())
			}
		case linebot.EventSourceTypeRoom:
			if err := this.replyText(replyToken, "Leaving room"); err != nil {
				return err
			}
			if _, err := this.bot.LeaveRoom(source.RoomID).Do(); err != nil {
				return this.replyText(replyToken, err.Error())
			}
		}
	default:
		//log.Printf("Echo message to %s: %s", replyToken, message.Text)
		//if err := this.replyText(replyToken, message.Text); err != nil {
		//return err
		//}
	}
	return nil
}

func (this *Bot) handleImage(message *linebot.ImageMessage, replyToken string) error {
	return this.handleHeavyContent(message.ID, func(content *os.File) error {
		contentUrl := this.BaseUrl + "/downloaded/" + filepath.Base(content.Name())
		if _, err := this.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(contentUrl),
			//linebot.NewImageMessage(contentUrl),
		).Do(); err != nil {
			return err
		}
		return nil
	})
}

func (this *Bot) handleHeavyContent(messageID string, callback func(*os.File) error) error {
	content, err := this.bot.GetMessageContent(messageID).Do()
	if err != nil {
		return err
	}
	defer content.Content.Close()
	log.Printf("Got file: %s", content.ContentType)
	originalConent, err := this.saveContent(content.Content)
	if err != nil {
		return err
	}
	return callback(originalConent)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (this *Bot) saveContent(content io.ReadCloser) (*os.File, error) {
	name := filepath.Join(this.DownloadDir, RandStringBytes(64))
	file, err := os.Create(name)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return nil, err
	}
	log.Printf("Saved %s", file.Name())
	return file, nil
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
				if err := this.handleImage(message, event.ReplyToken); err != nil {
					log.Print(err)
				}
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