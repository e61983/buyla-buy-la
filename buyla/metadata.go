package Buyla

type Group struct {
	IsOpen  bool               `json:"is_opening"`
	Records map[string]*Record `json:"records"`
}

type Record struct {
	UserProfile *Profile `json:"user_profile"`
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
	Size           string `json:"size"`
	Comment        string `json:"comment"`
	Id             int    `json:"id"`
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
	test_profile := NewProfile("星期天配音是對的", "https://purr.objects-us-east-1.dream.io/i/L5rgK.jpg")
	g.Records[test_uid] = NewRecord(test_profile)
	r := g.Records[test_uid]
	r.Goods = append(r.Goods, NewGood("國文", "無糖", "熱的", "1", "大", "在有跟沒有之間", 1654))
	r.Goods = append(r.Goods, NewGood("英文", "無糖", "熱的", "1", "大", "在有跟沒有之間", 1432))

	test_uid = "test-uid2"
	test_profile = NewProfile("星期天配音是不對的", "https://purr.objects-us-east-1.dream.io/i/KiX13.png")
	g.Records[test_uid] = NewRecord(test_profile)
	r = g.Records[test_uid]
	r.Goods = append(r.Goods, NewGood("法文", "少糖", "去冰", "1", "大", "在有跟沒有之間", 987651))
	r.Goods = append(r.Goods, NewGood("日文", "無糖", "熱的", "1", "大", "在有跟沒有之間", 9731))

	g.IsOpen = true
	return g
}

func NewProfile(name, photoUrl string) *Profile {
	return &Profile{DisplayName: name, PhotoUrl: photoUrl}
}

func NewRecord(profile *Profile) *Record {
	r := &Record{UserProfile: profile}
	r.Goods = make([]*Good, 0, 10)
	return r
}

func NewGood(itemName, sweetnessLevel, amountOfIce, number, size, comment string, id int) *Good {
	return &Good{
		ItemName:       itemName,
		SweetnessLevel: sweetnessLevel,
		AmountOfIce:    amountOfIce,
		Number:         number,
		Size:           size,
		Comment:        comment,
		Id:             id,
	}
}
