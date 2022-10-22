package category

import (
    "contract/pkg/app"
    "contract/pkg/database"
    "contract/pkg/paginator"
    "github.com/gin-gonic/gin"
)

func Get(idstr string) (categories Category) {
    database.DB.Where("id", idstr).First(&categories)
    return
}

func GetBy(field, value string) (categories Category) {
    database.DB.Where("? = ?", field, value).First(&categories)
    return
}

func All() (categories []Category) {
    database.DB.Find(&categories)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Category{}),
        &categories,
        app.V1URL(database.TableName(&Category{})),
        perPage,
    )
    return
}