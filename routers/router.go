package routers

import (
	"updatemanager/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/getboards", &controllers.MainController{}, "get:ChanelBoard")
    beego.Router("/addupdateapk", &controllers.MainController{}, "*:AddUpdateApk")
    beego.Router("/activatechargefee", &controllers.MainController{}, "*:ActivateChargefee")

    //charge balance entry
    beego.Router("/charge", &controllers.MainController{}, "get:Charge")
    beego.Router("/chargebalance", &controllers.MainController{}, "post:ChargeBalance")

    //query order
    beego.Router("/scanorder", &controllers.MainController{}, "post:ScanOrder")
    beego.Router("/findboss", &controllers.MainController{}, "post:FindBoss")

    //register chanelid
    beego.Router("/updatechanelid", &controllers.MainController{}, "post:UpdateChanelId")
   //首页
    beego.Router("/", &controllers.IndexController{}, "*:Index")
    //登录
    beego.Router("admin/login", &controllers.AccountController{}, "*:Login")
    //退出登陆
    beego.Router("admin/logout", &controllers.AccountController{}, "*:Logout")
    //修改个人资料
    beego.Router("admin/account/profile", &controllers.AccountController{}, "*:Profile")
    //添加用户
    beego.Router("admin/user/add", &controllers.CustomerUserController{}, "*:Add")
    //用户管理
    beego.Router("admin/user/list", &controllers.CustomerUserController{}, "*:List")//   /admin/user/list?page=3
    beego.Router("admin/user/edit", &controllers.CustomerUserController{}, "*:Edit")
    beego.Router("admin/user/delete", &controllers.CustomerUserController{}, "*:Delete")


}



