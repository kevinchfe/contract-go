package database

import (
	"contract/pkg/config"
	"database/sql"
	"errors"
	"fmt"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, _logger gormlogger.Interface) {
	// 使用gorm.open连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: _logger,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	// 获取底层的 sqlDB
	SQLDB, err = DB.DB()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func CurrentDatabase() (dbname string) {
	dbname = DB.Migrator().CurrentDatabase()
	return
}

func DeleteAllTables() error {
	var err error
	switch config.Get("database.connection") {
	case "mysql":
		err = deleteMysqlTables()
	case "sqlite":
		err = deleteAllSqliteTables()
	default:
		panic(errors.New("database connect not supported"))
	}
	return err
}

// TableName 获取表名称
func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	stmt.Parse(obj)
	return stmt.Schema.Table
}

func deleteAllSqliteTables() error {
	var tables []string
	// 读取所有表
	err := DB.Select(&tables, "SELECT name FROM sqlite_master WHERE type='table'").Error
	if err != nil {
		return err
	}
	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	return nil
}

func deleteMysqlTables() error {
	dbname := CurrentDatabase()
	tables := []string{}
	// 读取所有表
	err := DB.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	DB.Exec("SET foreign_key_checks = 0;")
	// 删除所有表
	for _, table := range tables {
		err := DB.Migrator().DropTable(table)
		if err != nil {
			return err
		}
	}
	// 开启外键检测
	DB.Exec("SET foreign_key_checks = 1;")
	return nil
}
