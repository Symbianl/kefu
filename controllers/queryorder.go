package controllers

import(
	"github.com/astaxie/beego"
	paysys2model "paysys2/models"
	"updatemanager/persistence"
	paysyschanel "paysys2/public/chanel"
	paysystools "paysys2/wechat/public/tools"
	paysysdb "paysys2/paysys/db"
	"github.com/chanxuehong/wechat/mch"
	"crypto/subtle"
	paysysutils "paysys2/public/utils"
	"time"
	"strconv"
)

func (this *MainController) ScanOrder() {
	this.TplName = "index.tpl"
	orderNo := this.Input().Get("no")
	beego.Debug("[SCANORDER] order no:", orderNo)
	order := persistence.GetOrderInfo(orderNo)
	if order == nil {
		beego.Debug("[SCANORDER] query new order not found, order no:", orderNo)
		this.Ctx.WriteString(`{"code":-1, "desc":"订单未找到！请核实订单号"}`)
		return
	}
	ok, result := QueryAndRecordOrder(order)
	if !ok {
		this.Ctx.WriteString(`{"code":-2, "desc":"查询补单失败，` + result + `"}`)
		return	
	}
	this.Ctx.WriteString(`{"code":0, "desc":"查询并补单成功"}`)
	return
}

func QueryAndRecordOrder(order *paysys2model.Order) (bool, string) {
	beego.Debug("[SCANORDER] query and record order:", order)
	ok, _ := CheckScanPayResult(order)
	if !ok {
		return false, "check failed"
	}
	return true, "ok"
}


func CheckScanPayResult(order *paysys2model.Order) (bool, string) {	
	chanelInfo, ok := paysyschanel.GetChanelInfoOfBoard(order.BoardId)
	if !ok {
		beego.Debug("[SCANPAY WECHAT]get chanel info for checking payment bill failed!")
		return false, ""
	}
	beego.Informational("[SCANPAY] chanel mchid:", chanelInfo.ChanelWxPaymentMchid, " chanel parentid:", chanelInfo.ChanelParentId)
	subChanel := new(paysys2model.ChanelConfig)
	if chanelInfo.ChanelParentId != 0 {
		beego.Debug("[SCANPAY WECHAT] chanel parent id:", chanelInfo.ChanelParentId)
		subChanel = chanelInfo
		chanelInfo, _ = paysyschanel.GetChanelInfoById(chanelInfo.ChanelParentId)
		if chanelInfo == nil {
			beego.Debug("[SCANPAY WECHAT] get parent chanel info err!")
			return false ,""
		}
		beego.Informational("[SCANPAY WECHAT] get parent chanel info:", chanelInfo.ChanelName)
	}
	beego.Informational("[SCANPAY WECHAT] chanel mchid:", chanelInfo.ChanelWxPaymentMchid, " sub chanel mchid:", subChanel.ChanelWxPaymentMchid)
	proxy := mch.NewProxy(chanelInfo.ChanelWxPaymentApikey, nil)
	paramLength := 5
	if subChanel.ChanelParentId != 0 {
		paramLength = 6
	}
	beego.Informational("[SCANPAY] parameters len:", paramLength)
	parameters := make(map[string]string, paramLength)
	if subChanel.ChanelParentId != 0 {
		parameters["sub_mch_id"] = subChanel.ChanelWxPaymentMchid //子支付渠道mchid
		beego.Informational("[SCANPAY WECHAT] sub mch id:", subChanel.ChanelWxPaymentMchid)
	}
	parameters["appid"] = chanelInfo.ChanelWxAppid
	parameters["mch_id"] = chanelInfo.ChanelWxPaymentMchid
	parameters["out_trade_no"] = order.OrderNo
	parameters["nonce_str"] = paysysutils.RandomDigitGenerator(12)

	signature := paysystools.Sign(parameters, chanelInfo.ChanelWxPaymentApikey)
	parameters["sign"] = signature
	beego.Informational("[SCANPAY WECHAT] checking order with parameters:", parameters)
	msg, err := proxy.PostXML("https://api.mch.weixin.qq.com/pay/orderquery", parameters)

	if err != nil {
		beego.Warning("[SCANPAY WECHAT]postxml err:", err.Error())
		return false, ""
	}
	beego.Informational("[SCANPAY WECHAT]feedback from wechat server for query order:", msg)
	returnCode, ok := msg["return_code"]
	if !ok || returnCode == mch.ReturnCodeSuccess { //两个条件有一个成立就表示成功返回
		beego.Informational("paysysy2", "checkorder", "wx:" , chanelInfo.ChanelId ,":",msg["openid"], msg["fee_type"], msg["transaction_id"], msg["is_subscribe"], msg["cash_fee"], msg["total_fee"], msg["trade_type"], msg["out_trade_no"] ,":" ,msg["trade_state"])
		beego.Debug("[SCANPAY WECHAT]payment return code success:", ok, returnCode)
		for k, v := range msg {
			beego.Critical("[SCANPAY WECHAT]key:", k, " value:", v)
		}
		if order.OrderNo != string(msg["out_trade_no"]) {
			beego.Critical("[SCANPAY WECHAT]<Checking Order>out_trade_no not fit")
			return false, ""
		}
		//校验appid
		haveAppId := msg["appid"]

		if len(haveAppId) != len(chanelInfo.ChanelWxAppid) {
			beego.Debug("[SCANPAY WECHAT]appid len not match!")
			beego.Critical("[SCANPAY WECHAT]<Checking Order> appid len not match!")
			return false, ""
		}
		if subtle.ConstantTimeCompare([]byte(haveAppId), []byte(chanelInfo.ChanelWxAppid)) != 1 {
			beego.Critical("[SCANPAY WECHAT]appid not match!", haveAppId, "|", chanelInfo.ChanelWxAppid)
			beego.Debug("[SCANPAY WECHAT]<Checking Order> appid not match!")
			return false, ""
		}

		haveMchId := msg["mch_id"]
		wantMchId := chanelInfo.ChanelWxPaymentMchid
		if len(haveMchId) != len(wantMchId) {
			beego.Critical("[SCANPAY WECHAT]chid len not match!", haveMchId, "||", wantMchId)
			beego.Debug("[SCANPAY WECHAT]<Checking Order> mchid len not match!")
			return false, ""
		}
		if subtle.ConstantTimeCompare([]byte(haveMchId), []byte(wantMchId)) != 1 {
			beego.Critical("[SCANPAY WECHAT]mchid not match:", haveMchId, "||", wantMchId)
			beego.Debug("[SCANPAY WECHAT]<Checking Order> mchid not match!")
			return false, ""
		}

		signature1, ok := msg["sign"]
		if !ok {
			beego.Debug("[SCANPAY WECHAT]No sign parameter received!")
			beego.Debug("[SCANPAY WECHAT]<Checking Order> No sign found in post from wechat!")
			return false, ""
		}
		signature2 := paysystools.Sign(msg, chanelInfo.ChanelWxPaymentApikey)
		if len(signature1) != len(signature2) {
			beego.Debug("[SCANPAY WECHAT]sign len not match!")
			beego.Debug("[SCANPAY WECHAT]<Checking Order> sign len not match!")
			return false, ""
		}
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			beego.Critical("[SCANPAY WECHAT]sign not match!", signature1, "||", signature2)
			beego.Debug("[SCANPAY WECHAT]<Checking Order> sign not match!")
			return false, ""
		}
		if msg["trade_state"] == "SUCCESS" && msg["result_code"] == "SUCCESS" && msg["return_code"] == "SUCCESS"{
			beego.Informational("[SCANPAY WECHAT] saving paied order:", msg)
			finishedOrder := new(paysys2model.Order)
			finishedOrder.OrderNo = order.OrderNo
			finishedOrder.BoardId = order.BoardId
			finishedOrder.ChanelId = chanelInfo.ChanelId
			finishedOrder.OpenId = msg["openid"]
			finishedOrder.Fee = order.Fee
			finishedOrder.Coins = order.Coins
			finishedOrder.Date = time.Now().Unix()
			payKey := order.PayKey
			finishedOrder.PayKey = payKey
			finishedOrder.CltOrderNo = order.CltOrderNo
			finishedOrder.WxBankType = msg["bank_type"]
			finishedOrder.WxFeeType = msg["fee_type"]
			finishedOrder.WxTransationId = msg["transaction_id"]
			finishedOrder.WxIsSubscribe = msg["is_subscribe"]
			cashFee, _ := strconv.ParseInt(msg["cash_fee"], 10 , 64)
			finishedOrder.WxCashFee = cashFee
			totalFee, _ := strconv.ParseInt(msg["total_fee"], 10, 64)
			finishedOrder.WxTotalFee = totalFee
			finishedOrder.WxTradeType = msg["trade_type"]
			finishedOrder.ScanPayCode = order.ScanPayCode
			go paysysdb.AddFinishedOrder(finishedOrder)			
			return true, "success"
		}		
		if msg["result_code"] == "SUCCESS" && msg["return_code"] == "SUCCESS" && msg["trade_state"] != "SUCCESS"{
			if msg["trade_state"] == "USERPAYING"{
				return true, "waiting"
			}
			if msg["trade_state"] == "NOTPAY" {
				return false, "notpay"
			}
			if msg["trade_state"] == "CLOSED" {
				return false, "closed"
			}
			if msg["trade_state"] == "REVOKED" {
				return false, "revoked"
			}
			if msg["trade_state"] ==  "PAYERROR" {
				return false, "payerror"
			}
		}
		return false, ""
	}	
	return false, ""
}