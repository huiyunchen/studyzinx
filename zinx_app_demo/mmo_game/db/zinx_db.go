package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//玩家对象
type ZinxDB struct {
	Conn map[int64]*gorm.DB
}

var ZinxDBObj *ZinxDB

func init() {
	ZinxDBObj = &ZinxDB{
		Conn: make(map[int64]*gorm.DB),
	}
}

func GetDB(serverID int64) *gorm.DB {
	if ZinxDBObj.Conn[serverID] != nil {
		fmt.Println(serverID, "我静态化了")
		return ZinxDBObj.Conn[serverID]
	}

	var dsn string
	if serverID == 1 {
		dsn = "root:chy@654123@(127.0.0.1:3306)/zinx?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		fmt.Println(serverID, "数据库连接成功")
		ZinxDBObj.Conn[serverID] = db
	} else if serverID == 2 {
		dsn = "root:chy@654123@(127.0.0.1:3306)/zinx2?charset=utf8mb4&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		fmt.Println(serverID, "数据库连接成功")
		ZinxDBObj.Conn[serverID] = db
	}

	return ZinxDBObj.Conn[serverID]
}
