package seeders

import (
    "fmt"
    "contract/database/factories"
    "contract/pkg/console"
    "contract/pkg/logger"
    "contract/pkg/seed"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedProjectsTable", func(db *gorm.DB) {

        projects  := factories.MakeProjects(10)

        result := db.Table("projects").Create(&projects)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}