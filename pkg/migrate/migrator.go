package migrate

import (
	"contract/pkg/console"
	"contract/pkg/database"
	"contract/pkg/file"
	"gorm.io/gorm"
	"os"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的 migrations 表里的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator 创建 Migrator 实例，用以执行迁移操作
func NewMigrator() *Migrator {
	// 初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	// migrations 不存在就创建
	migrator.createMigrationsTable()
	return migrator
}

// Up 执行所有未迁移过的文件
func (migrator *Migrator) Up() {
	// 读取所有迁移文件，确保按照时间顺序执行
	migrateFiles := migrator.readAllMigrationFiles()
	// 获取当前批次值
	batch := migrator.getBatch()
	// 获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)
	// 可以通过此值来判断数据库是否已是最新
	runed := false
	// 对迁移文件进行遍历，如果没有执行 就执行up操作
	for _, mfile := range migrateFiles {
		// 对比文件名称，看是否已经执行
		if mfile.isNotMigrated(migrations) {
			migrator.runUpMigration(mfile, batch)
			runed = true
		}
	}
	if !runed {
		console.Success("database is up to date.")
	}
}

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback() {
	// 获取最后一批次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch=?", lastMigration.Batch).Order("id DESC").Find(&migrations)
	// 回滚最有一批次迁移
	if !migrator.rollbackMigrations(migrations) {
		console.Success("[migrations] table is empty, nothing to rollback.")
	}
}

// 获取当前这个批次的值
func (migrator *Migrator) getBatch() int {
	batch := 1
	// 去最后执行的一条迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	// 如果有值的话 加一
	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

// 创建 migrations 表
func (migrator *Migrator) createMigrationsTable() {
	migration := Migration{}
	if !migrator.Migrator.HasTable(&migration) {
		migrator.Migrator.CreateTable(&migration)
	}
}

func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	// 读取 database/migrations/ 目录下的所有文件
	// 默认会按照文件名称进行排序
	files, err := os.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _, f := range files {
		// 去除文件后缀 .go
		fileName := file.FileNameWithoutExtension(f.Name())
		// 通过迁移文件的名称获取【MigrationFile】对象
		mfile := getMigrationFile(fileName)
		// 加个判断，确保迁移文件可用，再放进migrateFiles 数组中
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrateFiles, mfile)
		}
	}
	// 返回排序好的【MigrationFile】 数组
	return migrateFiles
}

// runUpMigration 执行迁移 up 方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile, batch int) {
	// 执行up区块的SQL
	if mfile.Up != nil {
		// 友好提示
		console.Warning("migrating " + mfile.FileName)
		// 执行up方法
		mfile.Up(database.DB.Migrator(), database.SQLDB)
		// 提示已迁移哪个文件
		console.Success("migrated " + mfile.FileName)
	}
	// 入库
	err := migrator.DB.Create(&Migration{Migration: mfile.FileName, Batch: batch}).Error
	console.ExitIf(err)
}

// rollbackMigrations 回退迁移，按照倒序执行迁移的 down 方法
func (migrator *Migrator) rollbackMigrations(migrations []Migration) bool {
	// 标记是否真的有执行了迁移回退操作
	runed := false

	for _, _migration := range migrations {
		console.Warning("rollback " + _migration.Migration)
		// 执行迁移文件 down 方法
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil {
			mfile.Down(database.DB.Migrator(), database.SQLDB)
		}
		runed = true
		// 回退成功了就删除掉这条记录
		migrator.DB.Delete(&_migration)
		console.Success("finish " + mfile.FileName)
	}
	return runed
}
