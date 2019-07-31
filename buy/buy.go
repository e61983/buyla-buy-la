package buy

import (
	"strings"
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
