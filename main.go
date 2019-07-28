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
	"buy"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client
var groups map[string]*buy.Group

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	if err != nil {
		log.Fatal("Line Bot", err)
	}

	http.HandleFunc("/callback", callbackHandler)

	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)

	groups = buy.NewGroups()

	http.ListenAndServe(addr, nil)
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

	var msg *linebot.TextMessage

	for _, event := range events {
		if event.Type == linebot.EventTypeMemberJoined || event.Type == linebot.EventTypeJoin {
			msgText := "Hello~~~~~\n"
			msgText += "大家可以試著用 \n開團, 我要XXX, 印出明細 以及 收單 關鍵字\n"
			msgText += "來揪團喔~!!\n"
			msg = linebot.NewTextMessage(msgText)
			if _, err = bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
				log.Print(err)
			}
			log.Println("type:", event.Type)
		}
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:

				log.Println("UserID", event.Source.UserID)
				log.Println("GroupID", event.Source.GroupID)
				log.Println("RoomID", event.Source.RoomID)

				if "" != event.Source.GroupID {
					if _, ok := groups[event.Source.GroupID]; ok {
						log.Println("God Group ID")
					} else {
						log.Println("Create New Group")
						groups[event.Source.GroupID] = buy.NewGroup(event.Source.GroupID)
					}
				} else {
				}

				if strings.Contains(message.Text, "開團") {
					if groups[event.Source.GroupID].IsOpening {
						msg = linebot.NewTextMessage("已經在開了喔~!")
					} else {
						msg = linebot.NewTextMessage("開團啦~~!!!!!\n以下開放下單\n--------------------- ")
						groups[event.Source.GroupID].IsOpening = true
						log.Println("IsOpening = ", groups[event.Source.GroupID].IsOpening)
					}
				}

				if strings.Contains(message.Text, "收單") {
					groups[event.Source.GroupID].IsOpening = false
					msg = linebot.NewTextMessage("收單!!!!!")
				}

				if strings.Contains(message.Text, "我要") {
					if groups[event.Source.GroupID].IsOpening {
						goods := strings.Replace(message.Text, "我要", "", 1)
						if _, ok := groups[event.Source.GroupID].Records[event.Source.UserID]; ok {
							res, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
							if err != nil {
								log.Println("GetProfile err:", err)
							}
							record := buy.NewRecord()
							record.UserName = res.DisplayName
							record.Goods = goods
							groups[event.Source.GroupID].Records[event.Source.UserID] = record
							log.Println("Modify Record - ", res.DisplayName)
						} else {
							res, err := bot.GetGroupMemberProfile(event.Source.GroupID, event.Source.UserID).Do()
							if err != nil {
								log.Println("GetProfile err:", err)
							}
							record := buy.NewRecord()
							record.UserName = res.DisplayName
							record.Goods = goods
							groups[event.Source.GroupID].Records[event.Source.UserID] = record
							log.Println("Modify Record - ", res.DisplayName)
						}
						msg = linebot.NewTextMessage("好喔~! " + groups[event.Source.GroupID].Records[event.Source.UserID].UserName + "要" + goods)
					}
				}

				if strings.Contains(message.Text, "印出明細") {
					msgText := "熱騰騰的明細出來啦~~\n"
					for _, record := range groups[event.Source.GroupID].Records {
						msgText = msgText + record.UserName + ": " + record.Goods
					}
					msg = linebot.NewTextMessage(msgText)
				}

			}
			if msg != nil {
				if _, err = bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
					log.Print(err)
				}
			}
		}
	}
}
