package models

import (
	"fmt"
	"log"
	"server/pkg/config"
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
	println(con.Driver)
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
	var currProduct Product
	var productList []Product
	res := Database.Table(ProductTableName).Where("deleted_at is null ").Find(&productList)
	if res.Error != nil || len(productList) == 0 {
		produt := Product{
			ProductName:   "小米11",
			ProductNumber: 50,
		}
		Database.Table(ProductTableName).Save(&produt)
		currProduct = produt
	} else {
		Database.Table(ProductTableName).First(&currProduct)
	}

	var secKillList []SecKill
	res = Database.Table(SecKillTableName).Where("deleted_at is null ").Find(&secKillList)
	h, _ := time.ParseDuration("1h")
	if res.Error != nil || len(secKillList) == 0 {
		seckill := SecKill{
			Product:   currProduct,
			StartTime: time.Now(),
			EndTime:   time.Now().Add(h),
		}
		Database.Table(SecKillTableName).Save(&seckill)
	}
}

func migration() {
	// 自动迁移模式
	err := Database.AutoMigrate(&Product{})
	if err != nil {
		log.Fatalf("AutoMigrate failed:%v", err)
	}

	err = Database.AutoMigrate(&Order{})
	if err != nil {
		log.Fatalf("AutoMigrate failed:%v", err)
	}

	err = Database.AutoMigrate(&SecKill{})
	if err != nil {
		log.Fatalf("AutoMigrate failed:%v", err)
	}
}
