package persistence

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	paysysmodels "paysys2/models"
)

func GetBoardInfos(boardId string) *paysysmodels.BoardInfos {
	var boardInfo *paysysmodels.BoardInfos 
	o := orm.NewOrm()
	err := o.Raw(`select * from board_infos where board_id=?`, boardId).QueryRow(&boardInfo)
	if err != nil {
		beego.Debug("[FIND BOSS] find boss board info err:", err.Error())
		return nil
	}
	beego.Debug("[FIND BOSS] find boss board info ok. board id:", boardId)
	return boardInfo
}

func GetDeliverBossInfo(deliverId int64) *paysysmodels.User {
	var boss *paysysmodels.User 
	o := orm.NewOrm()
	err := o.Raw(`select * from user where deliver_id=? and is_owner=1`, deliverId).QueryRow(&boss)
	if err != nil {
		beego.Debug("[FIND BOSS] find boss info err:", err.Error(), "deliverid:", deliverId)
		return nil
	}
	beego.Debug("[FIND BOSS] find boss info ok for deliver:", deliverId)
	return boss
} 