package controllers

import(
	"github.com/astaxie/beego"
	"updatemanager/persistence"
)



func (this *MainController) Charge(){
	this.TplName ="charge.html"
}

func (this *MainController) ChargeBalance(){
	this.TplName = "charge.html"
	partOrder := this.Input().Get("no")
	beego.Debug("part order no:", partOrder)
	order := persistence.GetOrderPaidInfo(partOrder)
	if order == nil {
		this.Ctx.WriteString(`{"code":-1, "desc":"订单未找到"}`)
		return
	}
	ok := persistence.ChargeFee(order)
	if !ok {
		this.Ctx.WriteString(`{"code":-2, "desc":"充值失败！请联系系统管理员"}`)
		return
	}
	this.Ctx.WriteString(`{"code":0, "desc":"充值成功，请通知玩家!"}`)
	return
}