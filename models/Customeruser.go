package models

import (

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

//用户
type CustomerUser struct {
	Id int		//主键
	Username 	string `orm:"size(15)"`    //用户名
	Password 	string `orm:"size(32)"`    	//密码
	Email 		string `orm:"size(50)"`       //邮箱
	Logincount int     //登陆次数
	Authkey string `orm:"size(10)"`    //用户key
	Active int       //是否激活
	status int64
}

//获取user结构体在数据库中所对应的表的表名
func (Customeruser *CustomerUser) TableName() string {
	dbprefix := beego.AppConfig.String("dbprefix")
	return dbprefix + "user"
}

//插入
func (Customeruser *CustomerUser) Insert() error {
	if _, err := orm.NewOrm().Insert(Customeruser); err != nil {
		return err
	}
	return nil
}

//读取
func (Customeruser *CustomerUser) Read(fields ...string) error {
	if err := orm.NewOrm().Read(Customeruser, fields...); err != nil {
		return err
	}
	return nil
}

//更新
func (Customeruser *CustomerUser) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(Customeruser, fields...); err != nil {
		return err
	}
	return nil
}

//删除
func (Customeruser *CustomerUser) Delete() error {
	if _, err := orm.NewOrm().Delete(Customeruser); err != nil {
		return err
	}
	return nil
}