package persistence

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"updatemanager/models"
)

func GetChanelBoards(userId int64) []*models.ChanelBoard {
	var boards []*models.ChanelBoard
	o := orm.NewOrm()
	_, err := o.Raw(`select * from board_infos where board_type=10`).QueryRows(&boards)	
	if err != nil {
		beego.Debug("[BOARD] get boards of chanel user err:", err.Error(), "userid:", userId)
		return nil
	}
	return boards
}

func AddBoardUpdateApk(boardId, url, apkName string) bool {
	o := orm.NewOrm()
	updateApk := GetBoardUpdateApk(boardId)
	if updateApk != nil {
		_, err := o.Raw(`update update_apk set apkname=?, url=? , updated=1 where board_id=?`, apkName, url, boardId).Exec()
		if err != nil {
			beego.Debug("[BOARD] update board apk err:", err.Error(), "boardid:", boardId)
			return false
		}
		beego.Debug("[BOARD] update new update apk info ok. boardid:", boardId)
		return true
	}
	_, err:= o.Raw(`insert into update_apk(board_id, url, apkname, updated) 
		values(?,?,?,1)`,boardId, url, apkName).Exec()
	if err != nil{
		beego.Debug("[BOARD] add new update apk info err:", err.Error(), "boardid:", boardId)
		return false
	}
	beego.Debug("[BOARD] add new update apk info ok. boardid:", boardId)
	return true
}

func GetBoardUpdateApk(boardId string) *models.UpdateApk {
	o := orm.NewOrm()
	var updateApk *models.UpdateApk
	err := o.Raw(`select * from update_apk where board_id=?`, boardId).QueryRow(&updateApk)
	if err != nil {
		beego.Debug("[UPDATE] get board update apk info err:", err.Error(), "boardid:", boardId)
		return nil
	}
	return updateApk
}