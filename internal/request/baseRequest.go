package requests

// Pagination 代表分页查询的基础参数
type Pagination struct {
	Page     int `json:"page" form:"page" binding:"gte=1"`     // 当前页数，最小为1
	PerPage  int `json:"per_page" form:"per_page" binding:"gte=1"` // 每页数量，最小为1
}

// ToMap 将Pagination转换为map，方便传递给数据库查询等操作
func (p Pagination) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"offset": (p.Page - 1) * p.PerPage,
		"limit":  p.PerPage,
	}
}