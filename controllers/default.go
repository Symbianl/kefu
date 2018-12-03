package controllers

import (
	"github.com/astaxie/beego"
	"updatemanager/persistence"
)

type MainController struct {
	baseController
}

func (this *MainController) ChanelBoard() {
	this.TplName = "update.html"
	boards := persistence.GetChanelBoards(38)
	this.Data["Boards"] = boards
}

func (this *MainController) AddUpdateApk(){
	boardId := this.Input().Get("bid")
	apkName := this.Input().Get("apkname")
	url := this.Input().Get("url")
	beego.Debug("[BOARD] add new update apk info, boardid:", boardId, " apkname:", apkName, "url:", url)
	ok := persistence.AddBoardUpdateApk(boardId, url, "lockers_dx.apk")
	if !ok {
		beego.Debug("[BOARD] add new update info failed.")
		this.Ctx.WriteString(`{"code":-1, "desc":"更新包配置失败，请联系技术支持"}`)
		return
	}
	beego.Debug("[BOARD] add new update info ok")
	this.Ctx.WriteString(`{"code":0, "desc":"更新包配置OK，请通知店家确保网络连接正常"}`)
	return
}

func (this *MainController) ActivateChargefee(){
	phone := this.Input().Get("phone")
	if phone == "" {
		this.Ctx.WriteString(`{"code":-1, "desc":"老板电话号码不能为空"}`)
		return
	}
	ok := persistence.ActivateChargefee(phone)
	if !ok {
		beego.Debug("[CHARGEFEE] activate chargefee failed.")
		this.Ctx.WriteString(`{"code":-2, "desc":"激活充值功能失败，请确认店家登记的手机号无误！"}`)
		return
	}
	beego.Debug("[CHARGEFEE] activate chargefee ok!")
	this.Ctx.WriteString(`{"code":0, "desc":"店家充值业务已经开启，请通知店家进行充值优惠配置"}`)
	return
}