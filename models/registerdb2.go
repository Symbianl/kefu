package models

import(
	"database/sql"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego"
)

const mysqlDataSource= "root:123456@tcp(127.0.0.1:3306)/mydb?charset=utf8"

var SqlDB *sql.DB

func init(){
	RegisterDB2()
}
func RegisterDB2(){
	var err error
	SqlDB, err = sql.Open("mysql", mysqlDataSource)
	if err != nil {
		beego.Debug("[INIT]init db connection err:", err.Error())
		os.Exit(-1)
	}
	if err = SqlDB.Ping(); err != nil {
		beego.Debug("[INIT] db ping err:", err.Error())
		os.Exit(-1)
	}
	beego.Debug("[INIT] db inition ok....sql db:", SqlDB)
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetMaxIdleConns(10)
}

