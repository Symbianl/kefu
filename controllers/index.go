package controllers

import(
	"github.com/astaxie/beego/orm"
	"updatemanager/models"
	)


type IndexController struct {
	baseController
}

func (this *IndexController) Index() {
	this.display()
	this.Data["usernum"], _ = orm.NewOrm().QueryTable(new(models.CustomerUser)).Count()
}