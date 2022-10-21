package paginator

import (
	"contract/pkg/config"
	"contract/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"strings"
)

// Paging 分页数据
type Paging struct {
	CurrentPage int
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageUrl string
	PrevPageUrl string
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL    string // 拼接url
	PerPage    int
	Page       int
	Offset     int // 数据库读取时的offset值
	TotalCount int64
	TotalPage  int
	Sort       string // 排序字段
	Order      string // 排序顺序

	query *gorm.DB     // db query 句柄
	ctx   *gin.Context // gin context
}

// Paginate 分页
// c gin.context 读取分页url参数
// db - Gorm 查询句柄
// baseURL - 分页链接
// data - 模型数组，传址获取数据
// 用法：
// query := database.DB.Model(Topic{}).Where("category_id=?", cid)
// var topics []Topic
// paging := paginator.Paginate(c, query,&topics,app.APIURL(database.TableName(&Topic{})),perPage)
func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	// 初始化 Paginator 实例
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseURL)
	// 查询数据库
	err := p.query.Preload(clause.Associations). // 读取关联
							Order(p.Sort + " " + p.Order). // 排序
							Limit(p.PerPage).
							Offset(p.Offset).
							Find(data).
							Error
	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}
	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageUrl: p.getNextPageURL(),
		PrevPageUrl: p.getPrevPageURL(),
	}
}

func (p *Paginator) initProperties(perPage int, baseURL string) {
	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	// 排序参数
	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

// 兼容 URL 带与不带`?`的情况
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

func (p *Paginator) getPerPage(perPage int) int {
	// 优先使用url请求中的per_page参数
	queryPerpage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerpage) > 0 {
		perPage = cast.ToInt(queryPerpage)
	}
	// 没有传参 使用默认
	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}
	return perPage
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

func (p *Paginator) getCurrentPage() int {
	// 优先使用url中的page参数
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page <= 0 {
		page = 1
	}
	// TotalPage 等于0
	if p.TotalPage == 0 {
		return 0
	}
	// 请求页数大于总页数
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

// getPageLink 拼接分页链接
func (p *Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

// getNextPageURL 下一页链接
func (p *Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

// getPrevPageURL 下一页链接
func (p *Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}
