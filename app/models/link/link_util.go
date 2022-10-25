package link

import (
    "contract/pkg/app"
    "contract/pkg/database"
    "contract/pkg/paginator"
    "github.com/gin-gonic/gin"
)

func Get(idstr string) (links Link) {
    database.DB.Where("id", idstr).First(&links)
    return
}

func GetBy(field, value string) (links Link) {
    database.DB.Where("? = ?", field, value).First(&links)
    return
}

func All() (links []Link) {
    database.DB.Find(&links)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Link{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Link{}),
        &links,
        app.V1URL(database.TableName(&Link{})),
        perPage,
    )
    return
}