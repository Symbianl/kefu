package controllers

import(
	"github.com/astaxie/beego"
	"updatemanager/models"
	paysys2models "paysys2/models"
	"encoding/json"
	"strconv"
)

func (this *MainController) CheckBoardPayInfo(){
	this.TplName = "index.tpl"
	boardId := this.Input().Get("id")
	beego.Debug("[CHECK BOARD] get boardid to check:",boardId)
	if models.RedisDB != nil {
		boardInfoStr, err := models.RedisDB.Get(boardId).Result()
		if err != nil {
			beego.Critical("[CHECK BOARD] get boardinfo in redis with boardid:", boardId , " err:", err.Error())	
			this.Ctx.WriteString(`{"code":-1, "desc":"boardinfo string not found"}`)
			return
		}else{
			beego.Critical("[CHECK BOARD] get boardinfo in redis with boardid ok:", boardId, " info:", boardInfoStr)
		}			
		var boardInfo *paysys2models.BoardInfos 
		json.Unmarshal([]byte(boardInfoStr), &boardInfo)
		if boardInfo == nil {
			beego.Critical("[CHECK BOARD] Unmarshal board info err. boardInfo nil, boardInfo str:", boardInfoStr)
			this.Ctx.WriteString(`{"code":-2, "desc":"board info unmashal failed"}`)
			return
		}
		chanelConfigStr, err := models.RedisDB.Get(strconv.FormatInt(boardInfo.ChanelId, 10)).Result()
		if err != nil {
			beego.Critical("[CHECK BOARD] get chanel info in redis with chanelid:", boardInfo.ChanelId, "err:", err.Error())
			this.Ctx.WriteString(`{"code":-3, "desc":"chanel info not found!"}`)
			return
		}
		beego.Critical("[CHECK BOARD] get boardinfo:", boardInfoStr, " chanelinfo:", chanelConfigStr)
		this.Ctx.WriteString(boardInfoStr + "<====>" + chanelConfigStr)
	}else{
		beego.Critical("[UPDATE CHANEL] redis db nil!!!!!")
	}		
}