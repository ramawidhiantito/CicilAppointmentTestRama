package main

import (
	//models "CicilAppointmentApp/models"
	_ "CicilAppointmentApp/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "database/orm_test.db")
	//orm.RegisterModel(new(models.Appointment))
}

func main() {
	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		beego.Error(err)
	}
	beego.Run()
}
