package models

import (
	"fmt"
	"log"
	"multi-cloud-poc/pkg/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	glogger "gorm.io/gorm/logger"
)

const mysqlDriverName = "mysql"
const sqliteDatabaseName = "sec_kill_db"

//Database database instance
var Database *gorm.DB

//Init init database instance
func Init() {
	con := config.GetDBConfig()
	switch con.Driver {
	case mysqlDriverName:
		mysqlDatabase(*con)
		break
	default:
		sqliteDatabase()
	}

}

// Connect to sqlite
func sqliteDatabase() {
	db, err := gorm.Open(sqlite.Open(sqliteDatabaseName), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		log.Fatalf("Connect sqlite err:%v", err)
	}
	Database = db
	migration()
}

// Connect to mysql
func mysqlDatabase(dbConfig config.DatabaseConfig) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Addr,
		dbConfig.DB)
	log.Printf("Connect to mysql %s", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: glogger.Default.LogMode(glogger.Silent),
	})
	if err != nil {
		log.Fatalf("Connect to mysql err:%v", err)
	}

	//2 Set MaxIdleConn/MaxOpenConn time
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("get sqlDB err:%v", err)
	}
	sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
	sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Second * time.Duration(dbConfig.ConnMaxLifeTime))
	Database = db

	//3 Migration db
	migration()
}

func migration() {
	// 自动迁移模式
	err := Database.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("AutoMigrate failed:%v", err)
	}
}
