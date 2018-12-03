package persistence

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"updatemanager/models"
)

func ActivateChargefee(phone string) bool {
	o := orm.NewOrm()
	var user *models.User
	err := o.Raw(`select * from user where phone=?`, phone).QueryRow(&user)
	if err != nil {
		beego.Debug("[CHARGEFEE] err:", err.Error(), " phone:", phone)
		return false
	}
	_, err = o.Raw(`update deliver set chargefee=1 where id=?`, user.DeliverId).Exec()
	if err != nil {
		beego.Debug("[CHARGEFEE] update chargefee err:", err.Error(), "deliver id:", user.DeliverId)
		return false
	}
	beego.Debug("[CHARGEFEE] activate chargefee ok. phone:", phone, " deliverid:", user.DeliverId)
	return true
}

func GetOrderPaidInfo(partNo string) *models.Order {
	if len(partNo) < 11 {
		beego.Debug("[CHARGEFEE] not enough order len. must bigger than 11:", partNo)
		return nil
	}
	var order *models.Order 
	o := orm.NewOrm()
	o.Using("query")
	err := o.Raw(`select * from finished_orders_state where order_no like "%` + partNo+ `"`).QueryRow(&order)	
	if err != nil {
		beego.Debug("[CHARGEFEE] no order info found in recent state. part no:", partNo)
		//return nil
	}else{
		return order 
	}
	err = o.Raw(`select * from finished_orders_state_history where order_no like "%` + partNo+ `"`).QueryRow(&order)
	if err != nil {
		beego.Debug("[CHARGEFEE] order info not found in history state...", partNo, "err:", err.Error())
		return nil
	}
	return order
}

func ChargeFee(order *models.Order) bool {
	oldBalance := int64(0)
	var boardInfo *models.BoardInfos
	o := orm.NewOrm()
	err := o.Raw("select * from board_infos where board_id=?", order.BoardId).QueryRow(&boardInfo)
	if err != nil {
		beego.Debug("[CHARGEFEE] get board info err:", err.Error(), "bordid:", order.BoardId)
		return false
	}

	err = o.Raw(`select balance from
	 privatepay_consumer_deliver_balance
	 where open_id=? and deliver_id=?`, order.OpenId, boardInfo.DeliverId).QueryRow(&oldBalance)
	if err != nil {
		beego.Debug("[CHARGEFEE] get old balance err:", err.Error(), " openid:", order.OpenId, boardInfo.DeliverId)
		_, er := o.Raw(`insert into privatepay_consumer_deliver_balance(open_id, deliver_id) values(?,?)`, order.OpenId, boardInfo.DeliverId).Exec()
		if er != nil {
			beego.Debug("[CHARGEFEE] insert user deliver account err:", er.Error(), "openid:", order.OpenId, boardInfo.DeliverId)
			return false
		}
	}
	newBalance := oldBalance + order.Coins 
	_, err = o.Raw(`update privatepay_consumer_deliver_balance 
		set balance = ? where open_id=? and deliver_id=?`, newBalance, order.OpenId, boardInfo.DeliverId).Exec()
	if err != nil {
		beego.Debug("[CHARGEFEE] update player balance err:", err.Error(), " new balance:", newBalance, "oldBalance:", oldBalance, "order coins:", order.Coins)
		return false
	}
	beego.Debug("[CHARGEFEE] update player balance ok. new balance:", newBalance, "oldBalance:", oldBalance, "order coins:", order.Coins)
	return true
}