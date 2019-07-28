package buy

type Group struct {
	ID        string
	IsOpening bool
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
