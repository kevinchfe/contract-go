package factories

import (
	"contract/app/models/link"
	"github.com/bxcodec/faker/v3"
)

func MakeLinks(count int) []link.Link {

	var objs []link.Link

	// 设置唯一性，如 Link 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		linksModel := link.Link{
			Name: faker.Username(),
			URL:  faker.URL(),
		}
		objs = append(objs, linksModel)
	}

	return objs
}
