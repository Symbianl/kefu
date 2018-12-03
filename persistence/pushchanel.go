package persistence

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"paysys2/models"
)

func GetChanelInfo(chanelId int64) *models.ChanelConfig {
	o := orm.NewOrm()
	var chanelInfo *models.ChanelConfig 
	err := o.Raw(`select * from chanel_config where chanel_id=?`, chanelId).QueryRow(&chanelInfo)
	if err != nil {
		beego.Debug("[UPDATE CHANEL] get chanel config info err:", err.Error(), " chanelid:", chanelId)
		return nil
	}
	return chanelInfo
}

func GetChanelIdBoards(chanelId int64) []*models.BoardInfos {
	o := orm.NewOrm()
	var boards []*models.BoardInfos
	_, err := o.Raw(`select * from board_infos where chanel_id=?`, chanelId).QueryRows(&boards)
	if err != nil {
		beego.Debug("[UPDATE CHANEL] get chanel boads infos err:", err.Error(), " chanelid:", chanelId)
		return nil
	}
	return boards
}