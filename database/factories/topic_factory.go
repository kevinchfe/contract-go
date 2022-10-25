package factories

import (
	"contract/app/models/topic"
	"github.com/bxcodec/faker/v3"
)

func MakeTopics(count int) []topic.Topic {

	var objs []topic.Topic

	// 设置唯一性，如 Topic 模型的某个字段需要唯一，即可取消注释
	// faker.SetGenerateUniqueValues(true)

	for i := 0; i < count; i++ {
		topicsModel := topic.Topic{
			Title:      faker.Sentence(),
			Body:       faker.Paragraph(),
			CategoryID: "4",
			UserID:     "1",
		}
		objs = append(objs, topicsModel)
	}

	return objs
}
