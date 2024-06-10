package response

type BaseResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // 数据字段，根据需要传入具体类型或nil
}

const (
	SuccessCode = 20000 // 成功状态码
)

// cluster类业务状态码 40100 开始
const (
	ClusterErrPOST int = iota + 40100
	ClusterErrDELETE
	ClusterErrPUT
	ClusterErrGet
	ClusterErrGetList

	// ... 可以继续添加其他集群相关的状态码
)

// ClusterErrMsg 定义一个map来存储状态码到状态信息的映射
var ClusterErrMsg = map[int]string{
	ClusterErrPOST:    "集群添加失败",
	ClusterErrDELETE:  "集群删除失败",
	ClusterErrPUT:     "集群修改失败",
	ClusterErrGet:     "集群查询失败",
	ClusterErrGetList: "集群列表查询失败",
	// ... 添加其他状态码和对应的状态信息
}

// namespace类业务状态码 40200 开始
const (
	NamespaceSuccess int = iota + 40200
	NamespaceFailed
	NamespaceFound
	NamespaceError
	// ... 可以继续添加其他集群相关的状态码
)

// server服务器内部状态码 50100 开始
const (
	ServerSuccess int = iota + 50100
	ServerNamespaceFailed
	ServerNamespaceFound
	ServerNamespaceError
	// ... 可以继续添加其他集群相关的状态码
)

// 定义状态信息常量
const ()

// SuccessResponse 创建一个成功的响应
func SuccessResponse(data interface{}) BaseResponse {
	return BaseResponse{
		Code:    SuccessCode,
		Message: "success",
		Data:    data,
	}
}

// ErrorResponse 创建一个错误响应
func ErrorResponse(code int, message string) BaseResponse {
	return BaseResponse{
		Code:    code,
		Message: message,
	}
}
