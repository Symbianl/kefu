package queryorder

import(
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	paysys2models "paysys2/models"
	"time"
	"updatemanager/controllers"
	//"updatemanager/models"
	//"database/sql"
)


func CheckingRecentNewOrders(){
	var orders []*paysys2models.Order 
	o := orm.NewOrm()
	_, err := o.Raw(`select * from new_orders_waiting where date >=?`, time.Now().Unix()-3000).QueryRows(&orders)
	if err != nil {
		beego.Debug("[UPDATE] no new orders found in table")
		return 
	}	
	for i :=0; i<len(orders); i++{
		beego.Debug("[UPDATE] checking order:", orders[i])
		ok, result := controllers.CheckScanPayResult(orders[i])
		beego.Debug("[UPDATE] checking order:", orders[i].OrderNo, " ok?", ok, "result:",result)
	}
}

/*
func CheckingRecentNewOrders(){
	if models.SqlDB == nil {
		beego.Debug("[UPDATE] sqldb nil!!!")
		models.RegisterDB2()
		beego.Debug("[UPDATE] sqldb after register:", models.SqlDB)
	}
	orders := []paysys2models.Order{}
	rows, err := models.SqlDB.Query("select * from new_orders_waiting where date >=?", time.Now().Unix()-300)
	if err != nil {
		beego.Debug("[UPDATE] get new orders err:", err.Error())
		return
	}
	defer rows.Close()
	for rows.Next(){
		order := paysys2models.Order{}
		err := rows.Scan(&order.OrderNo, &order.BoardId, &order.OpenId, &order.ProductId, &order.Fee, &order.Coins,
			&order.Date, &order.PrepayId, &order.State, &order.IsOnline, &order.WxBankType, 
			&order.WxFeeType, &order.WxTransationId, &order.WxIsSubscribe, &order.WxCashFee, &order.WxTotalFee,
			&order.WxTradeType, &order.PayKey, &order.CltOrderNo, &order.Discount, &order.ScanPayCode)
		if err != nil {
			beego.Debug("[UPDATE] scan order info in struct err:", err.Error())
		}else{
			orders = append(orders, order)
		}
	}
	for i := 0; i < len(orders); i++{
		ok, result := controllers.CheckScanPayResult(&orders[i])
		beego.Debug("[UPDATE] checking order:", orders[i].OrderNo, " ok?", ok, "result:",result)
	}
}

*/