package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	paysysmodels "paysys2/models"
	"github.com/go-redis/redis"

	"time"
	"crypto/md5"
	"fmt"
)

var(
	RedisDB *redis.Client
)

func init(){
	RegisterDB()
}
func RegisterDB() {
	maxIdle := 15
	maxConn := 15
	orm.RegisterModel(new(paysysmodels.ChanelConfig),new(CustomerUser))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/mydb?charset=utf8", maxIdle, maxConn)
	if err != nil {
		beego.Debug("default db:", err.Error())
	}

	RedisDB = redis.NewClient(&redis.Options{
		Addr:		"127.0.0.1:6379",
		Password: 	"123456",
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		DB:	0,
	})	
}

func Md5(buf []byte) string {
	mymd5 := md5.New()
	mymd5.Write(buf)
	result := mymd5.Sum(nil)
	return fmt.Sprintf("%x", result)
}

