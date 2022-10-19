package cmd

import (
	"contract/database/migrations"
	"contract/pkg/migrate"
	"github.com/spf13/cobra"
)

var CmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "Run database migration",
}

var CmdMigrateUp = &cobra.Command{
	Use:   "up",
	Short: "Run unmigrated migrations",
	Run:   runUp,
}

var CmdMigrateDown = &cobra.Command{
	Use:     "down",
	Aliases: []string{"rollback"}, // 别名 migrate down == migrate rollback
	Short:   "Reverse the up command",
	Run:     runDown,
}

func init() {
	CmdMigrate.AddCommand(
		CmdMigrateUp,
		CmdMigrateDown,
	)
}

func migrator() *migrate.Migrator {
	// 注册 database/migrations 下所有迁移文件
	migrations.Initialize()
	// 初始化migrator
	return migrate.NewMigrator()
}

func runUp(cmd *cobra.Command, args []string) {
	migrator().Up()
}

func runDown(cmd *cobra.Command, args []string) {
	migrator().Rollback()
}
