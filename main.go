package main

import (
	_ "updatemanager/routers"
	"github.com/astaxie/beego"
		"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
//	"updatemanager/models"
	"time"
	"updatemanager/queryorder"
	"net/http"
	"html/template"
)

func init(){	
	maxIdle := 15
	maxConn := 15
	err := orm.RegisterDataBase("query", "mysql", "root:123456@tcp(127.0.0.1:3306)/youbon_querys?charset=utf8", maxIdle, maxConn)
	if err != nil {
		beego.Debug("query db:", err.Error())
	}
	beego.Debug("[UPDATE] register query database ok.")	
}

func main() {
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//go queryorder.CheckingRecentNewOrders()
	go refreshQueryOrder()
	beego.ErrorHandler("404", page_not_found)
	beego.Run()
}

func page_not_found(rw http.ResponseWriter, r *http.Request) {
	t, _ := template.New("404.tpl").ParseFiles("views/404.tpl")

	data := make(map[string]interface{})
	data["content"] = "page not found"
	t.Execute(rw, data)
}

func refreshQueryOrder() {
	routine := time.NewTicker(300 * time.Second)
	beego.Debug("[UPDATE] refresh new orders data")
	for {
		select {
		case <-routine.C:
			queryorder.CheckingRecentNewOrders()			
		}
	}
}