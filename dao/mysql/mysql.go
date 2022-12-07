package mysql

import (
	"armor_plate/core"
	"armor_plate/model"
	"armor_plate/model/life"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init(cfg *core.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic("failed to open mysql")
		return
	}
	//设置最大连接数
	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	
	db.AutoMigrate(
		&model.Product{},
		&model.Employee{},
		&model.Category{},
		&model.Address{},
		&model.Order{},
		&model.Department{},
		&model.CasbinModel{},
		&model.UserAuthority{},
		&model.Depot{},
		&model.OutgoingOrder{},
		&life.Good{},
		&life.GoodCategory{},
		&life.Merchant{},
		&life.MerchantCategory{})
	return
}
