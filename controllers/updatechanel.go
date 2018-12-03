package controllers

import(
	"github.com/astaxie/beego"
	"updatemanager/persistence"
	"updatemanager/models"
	"strconv"
	"encoding/json"

)

func (this *MainController) UpdateChanelId(){
	this.TplName = "index.tpl"
	cId := this.Input().Get("id")
	beego.Debug("[UPDATE CHANEL] get chanelid to push:",cId)
	chanelId, _ := strconv.ParseInt(cId, 10, 64)
	if chanelId == 0 {
		beego.Debug("[UPDATE CHANEL] chanelid 0, just return")
		this.Ctx.WriteString(`{"code":-1, "desc":"渠道id非法"}`)
		return
	}	
	chanelInfo := persistence.GetChanelInfo(chanelId)
	if chanelInfo == nil {
		beego.Debug("[UPDATE CHANEL] chanelinfo nil.")
		this.Ctx.WriteString(`{"code":-2, "desc":"渠道信息为空！"}`)
		return
	}
	chanelInfoStr, _ := json.Marshal(chanelInfo)
	if models.RedisDB != nil {
		err := models.RedisDB.Set(strconv.FormatInt(chanelInfo.ChanelId,10), string(chanelInfoStr), 0).Err()
		if err != nil {
			beego.Critical("[UPDATE CHANEL] save chanelinfo in redis with key chanlid:", chanelInfo.ChanelId , " err:", err.Error())	
		}else{
			beego.Critical("[UPDATE CHANEL] save chanelinfo in redis with key chanelid ok, key:", chanelInfo.ChanelId)
		}	
		boards := persistence.GetChanelIdBoards(chanelId)
		for i := 0; i < len(boards); i++ {
			boardInfoStr, _ := json.Marshal(boards[i])
			err := models.RedisDB.Set(boards[i].BoardId, string(boardInfoStr), 0).Err()
			if err != nil {
				beego.Critical("[UPDATE CHANEL] save board in redis with key boardid:", boards[i].BoardId , " err:", err.Error())	
			}else{
				beego.Critical("[UPDATE CHANEL] save board in redis with key boardid ok, key:", boards[i].BoardId)
			}
		}
	}else{
		beego.Critical("[UPDATE CHANEL] redis db nil!!!!!")
	}		
}