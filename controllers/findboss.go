package controllers

import(
	"github.com/astaxie/beego"
	"updatemanager/persistence"
)

func (this *MainController) FindBoss(){
	beego.Debug("[FINDBOSS] finding boss...")
	this.TplName = "index.tpl"
	partOrder := this.Input().Get("no")
	beego.Debug("part order no:", partOrder)
	order := persistence.GetOrderPaidInfo(partOrder)
	if order == nil {
		this.Ctx.WriteString(`{"code":-1, "desc":"订单未找到"}`)
		return
	}
	boardInfo := persistence.GetBoardInfos(order.BoardId)
	if boardInfo == nil {
		this.Ctx.WriteString(`{"code":-2, "desc":"设备信息未找到！"}`)
		return
	}
	bossInfo := persistence.GetDeliverBossInfo(boardInfo.DeliverId)
	if boardInfo == nil {
		this.Ctx.WriteString(`{"code":-3, "desc":"店家老板信息未找！"}`)
		return
	}
	this.Ctx.WriteString(`{"code":0, "desc":"店家老板手机号:` + bossInfo.Phone + `"}`)
	return
}