package dto

// UserLoginRequest 用于获取用户登录的请求参数
type UserLoginRequest struct {
	Username string `json:"user_name" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"`  // 密码
}

// UserCreateRequest 用于获取用户创建的请求参数
type UserCreateRequest struct {
	Username string `json:"user_name" binding:"required"` // 用户名
	Password string `json:"password" binding:"required"`  // 密码
}
