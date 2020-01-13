package Buyla

type Group struct {
	IsOpen  bool               `json:"is_opening"`
	Records map[string]*Record `json:"records"`
}

type Record struct {
	UserProfile *Profile `json:"username"`
	Goods       []*Good  `json:"goods"`
	Comment     string   `json:"comment"`
}

type Profile struct {
	DisplayName string `json:"display_name"`
	PhotoUrl    string `json:"photo_url"`
}

type Good struct {
	ItemName       string `json:"item_name"`
	SweetnessLevel string `json:"sweetness_level"`
	AmountOfIce    string `json:"amount_of_ice"`
	Number         string `json:"number"`
}

type MetaData struct {
	Groups map[string]*Group `json:"data"`
}

func NewMetaData() *MetaData {
	m := &MetaData{}
	m.Groups = make(map[string]*Group)
	// TODO: just for test
	m.Groups["test"] = NewGroup()
	return m
}

func NewGroup() *Group {
	g := &Group{IsOpen: false}
	g.Records = make(map[string]*Record)
	// TODO: just for test
	test_uid := "test-uid"
	test_profile := NewProfile("星期天配音是對的", "https://randomuser.me/api/portraits/lego/5.jpg")
	g.Records[test_uid] = NewRecord(test_profile)
	g.IsOpen = true
	return g
}

func NewProfile(name, photoUrl string) *Profile {
	return &Profile{DisplayName: name, PhotoUrl: photoUrl}
}

func NewRecord(profile *Profile) *Record {
	r := &Record{UserProfile: profile}
	r.Goods = make([]*Good, 0, 10)
	// TODO: just for test
	r.Goods = append(r.Goods, NewGood("大平台", "大大", "大", "無限"))
	return r
}

func NewGood(itemName, sweetnessLevel, amountOfIce, number string) *Good {
	return &Good{ItemName: itemName, SweetnessLevel: sweetnessLevel, AmountOfIce: amountOfIce, Number: number}
}
