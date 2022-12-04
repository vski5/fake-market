package models

type GoodsCate struct {
	Id             int
	Title          string
	CateImg        string
	Link           string
	Template       string
	Pid            int
	SubTitle       string
	Keywords       string
	Description    string
	Sort           int
	Status         int
	AddTime        int
	GoodsCateItems []GoodsCate `gorm:"foreignKey:pid;references:Id"`
}

func (GoodsCate) TableName() string {
	return "goods_cate"
}

type GoodsType struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}

func (GoodsType) TableName() string {
	return "goods_type"
}

// 商品类型属性
type GoodsTypeAttribute struct {
	Id        int
	CateId    int
	Title     string
	AttrType  int
	AttrValue string
	Status    int
	Sort      int
	AddTime   int
}

func (GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
