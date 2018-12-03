package persistence

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	paysys2model "paysys2/models"
)

func GetOrderInfo(orderNo string) *paysys2model.Order {
	o := orm.NewOrm()
	var order *paysys2model.Order 
	err := o.Raw(`select * from new_orders where order_no=?`, orderNo).QueryRow(&order)
	if err != nil {
		beego.Debug("[QUERY] get new order infos in recent err:", err.Error(), "order no:", orderNo)
	}else{
		beego.Debug("[QUERY] get new order from recent. order:", order)
		return order
	}
	err = o.Raw(`select * from new_orders_history where order_no=?`, orderNo).QueryRow(&order)
	if err != nil {
		beego.Debug("[QUERY] get new order infos err:", err.Error(), "order no:", orderNo)
		return nil
	}
	beego.Debug("[QUERY] found new order info:", order)
	return order
} 