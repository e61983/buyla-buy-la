package Buyla

type OrderItem struct {
	ItemName       string `json:"goods"`
	SweetnessLevel string `json:"sweetness_level"`
	AmountOfIce    string `json:"amount_of_ice"`
	Number         string `json:"number"`
}

type OrderItems struct {
	List []*OrderItem `json:"list"`
}

type Record struct {
	UserName string      `json:"username"`
	Order    *OrderItems `json:"order"`
}

type Group struct {
	IsOpen  bool               `json:"is_opening"`
	Records map[string]*Record `json:"records"`
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
	return g
}

func NewRecord(name string) *Record {
	r := &Record{UserName: name}
	r.Order = NewOrderItems()
	return r
}

func NewOrderItems() *OrderItems {
	i := &OrderItems{}
	i.List = make([]*OrderItem, 0, 10)
	return i
}

func NewOrderItem(itemName, sweetnessLevel, amountOfIce, number string) *OrderItem {
	return &OrderItem{ItemName: itemName, SweetnessLevel: sweetnessLevel, AmountOfIce: amountOfIce, Number: number}
}
