package topic

import (
    "contract/pkg/app"
    "contract/pkg/database"
    "contract/pkg/paginator"
    "github.com/gin-gonic/gin"
)

func Get(idstr string) (topics Topic) {
    database.DB.Where("id", idstr).First(&topics)
    return
}

func GetBy(field, value string) (topics Topic) {
    database.DB.Where("? = ?", field, value).First(&topics)
    return
}

func All() (topics []Topic) {
    database.DB.Find(&topics)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Topic{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Topic{}),
        &topics,
        app.V1URL(database.TableName(&Topic{})),
        perPage,
    )
    return
}