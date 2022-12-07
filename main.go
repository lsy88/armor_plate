package main

import (
	"armor_plate/core"
	"armor_plate/dao/mysql"
	"armor_plate/pkg/snowflake"
	"armor_plate/router"
	"fmt"
)

func main() {
	//加载配置
	if err := core.Init(); err != nil {
		fmt.Printf("load config failed, err:#{err}\n")
		return
	}
	
	if err := core.ZapInit(core.Conf.LogConfig, core.Conf.Mode); err != nil {
		fmt.Printf("load config failed, err:#{err}\n")
		return
	}
	if err := mysql.Init(core.Conf.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	
	if err := snowflake.Init(1); err != nil {
		fmt.Printf("init sonwflake failed, err:#{err}\n")
		return
	}
	
	r := router.SetupRouter(core.Conf.Mode)
	err := r.Run(fmt.Sprintf(":%d", core.Conf.Port))
	if err != nil {
		fmt.Printf("run server failed, err:%v\n", err)
		return
	}
}
