package config

import "contract/pkg/config"

func init() {
	config.Add("paging", func() map[string]interface{} {
		return map[string]interface{}{
			// 默认每页条数
			"perpage": 10,
			// URL中用以分辨多少页的参数
			"url_query_page": "page",
			// URL中排序参数（如id）
			"url_query_sort": "sort",
			// URL中排序规则参数（正序或倒序）
			"url_query_order": "order",
			// URL中每页条数
			"url_query_per_page": "per_page",
		}
	})
}
