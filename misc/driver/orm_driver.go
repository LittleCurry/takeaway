package driver

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
)

var engine *xorm.Engine

func OrmInit(dsn string) {
	var err error
	engine, err = xorm.NewEngine("mysql", dsn)
	fmt.Println("engine:", engine)
	if err != nil {
		fmt.Println("OrmInit_err:", err)
		panic(err)
	}
	engine.ShowSQL(true)
}

func MySQL() *xorm.Engine {
	return engine
}