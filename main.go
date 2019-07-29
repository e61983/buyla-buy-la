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

func GetUsageString() string {
	msgText := "Hi~~~我是揪團啦\n"
	msgText += "大家可以試著用下面的幾個關鍵字來揪團喔~\n"
	msgText += "━ [開團]: 告訴大家有新的揪團!\n"
	msgText += "━ [我要]xxx: xxx 是你想訂的東西喔!\n"
	msgText += "━ [結單]: 就是告訴大家下回請早的意思啦~\n"
	msgText += "━ [明細]: 看看大家訂了什麼\n"
	msgText += "━ [說明]: 跟大家再自我介紹一次\n"
	msgText += "━ 叫你們 RD 出來滴霸格!!!: 沒什麼作用~只是發洩一下\n"

	return msgText
}

func EventTypeJoinHandler(event *linebot.Event) {
	msgText := GetUsageString()
	msg := linebot.NewTextMessage(msgText)
	if _, err := bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
		log.Print(err)
	}
}

func EventTypeMemberJoinedHandler(event *linebot.Event) {
	msgText := GetUsageString()
	msg := linebot.NewTextMessage(msgText)
	if _, err := bot.ReplyMessage(event.ReplyToken, msg).Do(); err != nil {
		log.Print(err)
	}
}

func GetAllRecordsString(groupID string) string {
	var msgText string
	recordNumber := len(groups[groupID].Records)
	if recordNumber == 0 {
		msgText = "好像..什麼也沒有喔~~  ˊ_>ˋ "
	} else {
		for _, record := range groups[groupID].Records {
			msgText = msgText + "━ " + record.UserName + " 要:\n " + record.Goods + "\n"
		}
	}

	return msgText
}

func EventTypeMessage_TextMessageHander(event *linebot.Event) {

	message := event.Message.(*linebot.TextMessage)
	var msg linebot.SendingMessage

	groupID := event.Source.GroupID
	userID := event.Source.UserID

	message.Text = strings.Replace(message.Text, "［", "[", 1)
	message.Text = strings.Replace(message.Text, "］", "]", 1)

	switch {
	case strings.Contains(message.Text, "[開團]"):
		if groups[groupID].IsOpening {
			msg = linebot.NewTextMessage("已經在開了喔~!")
		} else {
			groups[groupID].IsOpening = true
			groups[groupID].Records = buy.NewRecords()
			msg = linebot.NewTextMessage("開團啦~~!!!!!\n以下開放下單\n--------------------- ")
			log.Println("IsOpening = ", groups[groupID].IsOpening)
		}
	case strings.Contains(message.Text, "[結單]"):
		if groups[groupID].IsOpening {
			groups[groupID].IsOpening = false
			msg = linebot.NewTextMessage("結單啦!!!!! \n" + GetAllRecordsString(groupID))
			log.Println("IsOpening = ", groups[groupID].IsOpening)
		} else {
			msg = linebot.NewTextMessage("現在還沒有開始揪團~\n 大家都在等你開喔~!! XD")
		}
	case strings.Contains(message.Text, "[我要]"):
		if groups[groupID].IsOpening {
			goods := strings.Replace(message.Text, "[我要]", "", 1)
			res, err := bot.GetGroupMemberProfile(groupID, userID).Do()
			if err != nil {
				log.Println("GetProfile err:", err)
			}
			record := buy.NewRecord()
			record.UserName = res.DisplayName
			record.Goods = goods
			groups[groupID].Records[userID] = record
			log.Println("Modify Record - ", res.DisplayName)
			msg = linebot.NewTextMessage("好喔~! " + groups[groupID].Records[userID].UserName + "要" + goods)
		}
	case strings.Contains(message.Text, "[明細]"):
		msgText := "熱騰騰的明細出來啦~~\n"
		msgText += GetAllRecordsString(groupID)
		msg = linebot.NewTextMessage(msgText)
	case strings.Contains(message.Text, "[說明]"):
		msg = linebot.NewTextMessage(GetUsageString())
	case strings.Contains(message.Text, "叫你們 RD 出來滴霸格!!!"):
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
			EventTypeJoinHandler(event)
		case linebot.EventTypeJoin:
			EventTypeMemberJoinedHandler(event)
		case linebot.EventTypeMessage:
			switch event.Message.(type) {
			case *linebot.TextMessage:
				EventTypeMessage_TextMessageHander(event)
			}
		}
	}
}
