package requests

// GetUserListRequest 用于获取用户列表的请求参数
type GetUserListRequest struct {
	Pagination        // 嵌入分页结构体
	SearchTerm        string `json:"search_term" form:"search_term"` // 可选的搜索关键词
}

// GetUserRequest 用于获取单个用户的请求参数
type GetUserRequest struct {
	ID string `uri:"id" binding:"required"` // 用户ID，作为路径参数
}