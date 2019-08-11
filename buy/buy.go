package buy

import (
	"log"
	"net/http"
	"strings"
	"time"
)

type Group struct {
	ID        string
	IsOpening bool
	Store     string
	Records   map[string]*Record
}

type Record struct {
	UserName string
	Goods    string
}

func NewGroups() map[string]*Group {
	groups := make(map[string]*Group)
	return groups
}

func NewRecords() map[string]*Record {
	records := make(map[string]*Record)
	return records
}

func NewGroup(groupID string) *Group {
	group := &Group{ID: groupID, IsOpening: false}
	group.Records = NewRecords()
	return group
}

func NewRecord() *Record {
	record := &Record{UserName: "", Goods: ""}
	return record
}

func (g *Group) GetRecord(userName string) *Record {
	for key, record := range g.Records {
		if strings.EqualFold(record.UserName, userName) {
			return g.Records[key]
		}
	}
	return nil
}

func (g *Group) String() string {
	var msgText string
	recordNumber := len(g.Records)
	if recordNumber == 0 {
		msgText = "好像..什麼也沒有喔~~  ˊ_>ˋ "
	} else {
		for _, record := range g.Records {
			msgText = msgText + "━ " + record.UserName + " 要:\n " + record.Goods + "\n"
		}
	}
	return msgText
}

func (g *Group) AddUserGoods(userID, displayName, goods string) string {
	record := NewRecord()
	record.UserName = displayName
	record.Goods = goods
	g.Records[userID] = record

	return g.Records[userID].UserName + " 要\n" + goods
}

func (g *Group) RemoveUserGoods(userID string) {
	g.Records[userID] = nil
	delete(g.Records, userID)
	log.Println("Delete")
}

func SelfPing(url string) chan bool {
	log.Println("Enable self ping:" + url)
	return SetInterval(func() {
		resp, err := http.Get(url)
		log.Println("ping: Sending heartbeat to " + url)
		if err != nil {
			log.Printf("heroku-self-ping: Sending heartbeat error %s", err)
		}
		defer resp.Body.Close()
	}, 300000, false)
}

func SetInterval(doFunc func(), milliseconds int, async bool) chan bool {

	interval := time.Duration(milliseconds) * time.Millisecond

	ticker := time.NewTicker(interval)
	clear := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				if async {
					go doFunc()
				} else {
					doFunc()
				}
			case <-clear:
				log.Println("Disable self ping")
				ticker.Stop()
				return
			}
		}
	}()
	return clear

}
