package models

type ChanelBoard struct {
	BoardId 	string
	UserId 		int64
}

type UpdateApk struct {
	BoardId 	string 
	Updated 	int64
	Apkname 	string 
	Url 		string
}

type BoardInfos struct{
	BoardId       string `orm:"pk"`
	BoardIdShort  string
	BoardType     int64 `orm:"null"`
	BoardTypeName string
	SimPhone      string `orm:"null"`
	Location      string `orm:"null"`
	Manufacturer  string `orm:"null"`
	BatchNo       string `orm:"null"`
	Description   string `orm:"null"`
	Date          int64  `orm:"null"`
	Active        int64
	ChanelId      int64 `orm:"0"`
	ChanelName    string
	BoardName     string `orm:"null"`
	DeliverId     int64
	StockState    int64
	BoardOffline  int64
	Ready         int64
	State         string
	Transferer     int64
	ShopId        int64
}