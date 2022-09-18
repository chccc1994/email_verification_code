package models

import (
	"email_code/utils"
	"fmt"
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/go-redis/redis/v8"
)

// 初始化数据库
var DB *gorm.DB

var err error

func InitDb(){
	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName,
	)
	// fmt.Println(dns)
	DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	
	if err!=nil{
		log.Panicln("Database configuration error,",err.Error())
		return
	}
	_ = DB.AutoMigrate(&User{})

	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)
}

// Redis数据库初始化
func InitRedisDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}