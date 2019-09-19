// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"github.com/e61983/buyla-buy-la/buy"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var groups map[string]*buy.Group
var selfPing bool
var appName string
var selfPingFinish chan bool

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Fatal("Line Bot", err)
	}

	http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	appName = os.Getenv("HEROKU_APP_NAME")

	log.Println(appName)
	selfPing = false

	groups = buy.NewGroups()

	http.ListenAndServe(addr, nil)
}

func GetUsageString() string {
	msgText := "可以試著用下面的幾個關鍵字來揪團喔~\n"
	msgText += "- [開團]XXX : \n    告訴大家有新的揪團! 是要訂 XXX\n"
	msgText += "- [我要]xxx: \n    xxx 是你想訂的東西喔!\n"
	msgText += "- [我不要了]: \n    取消你原本相訂的東西\n"
	msgText += "- [加訂]xxx: \n    xxx 是你想加訂的東西!\n"
	msgText += "- [結單]: \n    就是告訴大家下回請早的意思啦~\n"
	msgText += "- [明細]: \n    看看大家訂了什麼\n"
	msgText += "- @XXX[咪兔]: \n    跟XXX 訂一樣的\n"
	msgText += "- [說明]: \n    跟大家再自我介紹一次\n"
	msgText += "- 叫你們 RD 出來滴霸格!!!: \n    沒什麼作用~只是發洩一下\n"

	return msgText
}

func EventTypeMemberJoinedHandler(event *linebot.Event) {
	var displayName string
	groupID := event.Source.GroupID
	userID := event.Source.UserID
	res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
	if err != nil {
		log.Println("GetProfile err:", err)
	} else {
		displayName = res.DisplayName
	}
	msgText := "Hi~~~" + displayName + "\n"
	msgText += GetUsageString()
	msg := linebot.NewTextMessage(msgText)
	if _, err := bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
		log.Print(err)
	}
}

func EventTypeJoinHandler(event *linebot.Event) {
	msgText := "Hi~~~我是揪團啦\n"
	msgText += GetUsageString()
	msg := linebot.NewTextMessage(msgText)
	if _, err := bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
		log.Print(err)
	}
}

func EventTypeMessage_TextMessageHander(event *linebot.Event) {

	message := event.Message.(*linebot.TextMessage)
	var msg linebot.SendingMessage

	groupID := event.Source.GroupID
	userID := event.Source.UserID
	currentGroup := groups[groupID]

	message.Text = strings.ReplaceAll(message.Text, "［", "[")
	message.Text = strings.ReplaceAll(message.Text, "］", "]")

	command, err := buy.ParseCommand(userID, message.Text)

	if err != nil {
		return
	}

	switch command.(type) {
	case *buy.OpenNewBuyLaCommand:
		if currentGroup.IsOpening {
			msg = linebot.NewTextMessage("已經在開了喔~!")
		} else {
			c := command.(*buy.OpenNewBuyLaCommand)
			if c.ShopName != "" {
				currentGroup.Store = c.ShopName
				currentGroup.IsOpening = true
				currentGroup.Records = buy.NewRecords()
				msg = linebot.NewTextMessage("開團啦~~!!!!!\n這次是 " + currentGroup.Store + " 喔\n\n------以下開放下單------\n ")
				log.Println("IsOpening = ", currentGroup.IsOpening)
			} else {
				currentGroup.Store = c.ShopName
				currentGroup.IsOpening = true
				currentGroup.Records = buy.NewRecords()
				msg = linebot.NewTextMessage("開團啦~~!!!!!\n\n------以下開放下單------\n ")
				log.Println("IsOpening = ", currentGroup.IsOpening)
			}
			if !selfPing {
				selfPingFinish = buy.SelfPing("https://" + appName + ".herokuapp.com")
				selfPing = true
			}
		}
	case *buy.CloseBuyLaCommand:
		if currentGroup.IsOpening {
			currentGroup.IsOpening = false
			msg = linebot.NewTextMessage("結單啦!!!!! \n" + currentGroup.String())
			log.Println("IsOpening = ", currentGroup.IsOpening)

			if selfPing {
				selfPingFinish <- true
				selfPing = false
			}

		} else {
			msg = linebot.NewTextMessage("現在還沒有開始揪團~\n 大家都在等你開喔~!! XD")
		}
	case *buy.WantCommand:
		if currentGroup.IsOpening {
			c := command.(*buy.WantCommand)
			res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
			if err != nil {
				log.Println("GetProfile err:", err)
			}
			msg = linebot.NewTextMessage("好喔~! " + currentGroup.AddUserGoods(userID, res.DisplayName, c.Goods))
		} else {
			msg = linebot.NewTextMessage("前一次揪團已結單\n等你開新團啦!")
		}
	case *buy.ShowRecordCommand:
		msgText := "熱騰騰的明細出來啦~~\n"
		msgText += currentGroup.String()
		msg = linebot.NewTextMessage(msgText)
	case *buy.CancelCommand:
		if currentGroup.IsOpening {
			record := currentGroup.Records[userID]
			msg = linebot.NewTextMessage("Okay~" + record.UserName + " 的\n" + record.Goods + "已經取消了")
			currentGroup.RemoveUserGoods(userID)
		} else {
			msg = linebot.NewTextMessage("目前還沒有開團喔")
		}
	case *buy.HelpCommand:
		msg = linebot.NewTextMessage(GetUsageString())
	case *buy.AttchCommand:
		if currentGroup.IsOpening {
			c := command.(*buy.AttchCommand)
			record := currentGroup.Records[userID]
			subCommand, err := buy.ParseCommand(userID, c.Goods)
			if err == nil {
				if sc, ok := subCommand.(*buy.MeTooCommand); ok {
					targetRecord := currentGroup.GetRecord(sc.TargetName)
					if targetRecord != nil && targetRecord.Goods != "" {
						res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
						if err != nil {
							log.Println("GetProfile err:", err)
						}
						msg = linebot.NewTextMessage("好喔~! " + currentGroup.AddUserGoods(userID, res.DisplayName, record.Goods+"\n"+targetRecord.Goods))
					} else {
						msg = linebot.NewTextMessage(sc.TargetName + " 還沒有訂喔!!!")
					}
				}
			} else {
				if record == nil {
					res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
					if err != nil {
						log.Println("GetProfile err:", err)
					}
					msg = linebot.NewTextMessage("好喔~! " + currentGroup.AddUserGoods(userID, res.DisplayName, c.Goods))
				} else {
					msg = linebot.NewTextMessage("好喔~! " + currentGroup.AddUserGoods(userID, record.UserName, record.Goods+"\n"+c.Goods))
				}
			}
		} else {
			msg = linebot.NewTextMessage("前一次揪團已結單\n等你開新團啦!")
		}
	case *buy.MeTooCommand:
		if currentGroup.IsOpening {
			c := command.(*buy.MeTooCommand)
			record := currentGroup.GetRecord(c.TargetName)
			if record != nil && record.Goods != "" {
				res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
				if err != nil {
					log.Println("GetProfile err:", err)
				}
				msg = linebot.NewTextMessage("好喔~! " + currentGroup.AddUserGoods(userID, res.DisplayName, record.Goods))
			} else {
				msg = linebot.NewTextMessage(c.TargetName + " 還沒有訂喔!!!")
			}
		} else {
			msg = linebot.NewTextMessage("前一次揪團已結單\n等你開新團啦!")
		}
	case *buy.RDDebugCommand:
		msg = linebot.NewStickerMessage("11537", "52002739")
	default:
	}

	if msg != nil {
		if _, err := bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
			log.Print(err)
		}
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {

	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {

		groupID := event.Source.GroupID
		if groupID == "" {
			return
		} else if _, ok := groups[groupID]; !ok {
			log.Println("Create a New Group")
			groups[groupID] = buy.NewGroup(groupID)
		}

		switch event.Type {
		case linebot.EventTypeMemberJoined:
			EventTypeMemberJoinedHandler(event)
		case linebot.EventTypeJoin:
			EventTypeJoinHandler(event)
		case linebot.EventTypeMessage:
			switch event.Message.(type) {
			case *linebot.TextMessage:
				EventTypeMessage_TextMessageHander(event)
			}
		}
	}
}
