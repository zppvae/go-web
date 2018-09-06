package main

import (
	_ "go-web/routers"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"go-web/models"

	"github.com/astaxie/beego"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func init() {
	models.RegisterDB()
}


func main() {
	// 开启 ORM 调试模式
	orm.Debug = true
	// create table
	orm.RunSyncdb("default", false, true)


	c, err := redis.Dial("tcp", "172.18.7.172:6379",redis.DialDatabase(8))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	_, err = c.Do("SET", "mykey", "superWang")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}


	// 启动 beego
	beego.Run()
}

