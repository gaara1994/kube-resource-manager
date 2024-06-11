package response

const (
	SuccessCode = 20000 // 成功状态码
)

// cluster类业务状态码 40100 开始
const (
	ClusterErrPost int = iota + 40100
	ClusterErrDelete
	ClusterErrPut
	ClusterErrGet
	ClusterErrGetList

	// ... 可以继续添加其他集群相关的状态码
)

// ClusterErrMsg 定义一个map来存储状态码到状态信息的映射
var ClusterErrMsg = map[int]string{
	ClusterErrPost:    "集群添加失败",
	ClusterErrDelete:  "集群删除失败",
	ClusterErrPut:     "集群修改失败",
	ClusterErrGet:     "集群查询失败",
	ClusterErrGetList: "集群列表查询失败",
	// ... 添加其他状态码和对应的状态信息
}

// namespace类业务状态码 40200 开始
const (
	NamespaceErrPost int = iota + 40200
	NamespaceErrDelete
	NamespaceErrPut
	NamespaceErrGet
	NamespaceErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var NamespaceErrMsg = map[int]string{
	NamespaceErrPost:    "名称空间添加失败",
	NamespaceErrDelete:  "名称空间删除失败",
	NamespaceErrPut:     "名称空间修改失败",
	NamespaceErrGet:     "名称空间查询失败",
	NamespaceErrGetList: "名称空间列表查询失败",
}

// 资源类型类业务状态码 40300 开始
const (
	ResourceTypeErrPost int = iota + 40300
	ResourceTypeErrDelete
	ResourceTypeErrPut
	ResourceTypeErrGet
	ResourceTypeErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var ResourceTypeErrMsg = map[int]string{
	ResourceTypeErrPost:    "资源类型添加失败",
	ResourceTypeErrDelete:  "资源类型删除失败",
	ResourceTypeErrPut:     "资源类型修改失败",
	ResourceTypeErrGet:     "资源类型查询失败",
	ResourceTypeErrGetList: "资源类型列表查询失败",
}

// 资源类业务状态码 40400 开始
const (
	ResourceConfigErrPost int = iota + 40400
	ResourceConfigErrDelete
	ResourceConfigErrPut
	ResourceConfigErrGet
	ResourceConfigErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var ResourceConfigErrMsg = map[int]string{
	ResourceConfigErrPost:    "资源类型添加失败",
	ResourceConfigErrDelete:  "资源类型删除失败",
	ResourceConfigErrPut:     "资源类型修改失败",
	ResourceConfigErrGet:     "资源类型查询失败",
	ResourceConfigErrGetList: "资源类型列表查询失败Config",
}

// 用户类业务状态码 40400 开始
const (
	UserErrPost int = iota + 40400
	UserErrDelete
	UserErrPut
	UserErrGet
	UserErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var UserErrMsg = map[int]string{
	UserErrPost:    "用户添加失败",
	UserErrDelete:  "用户删除失败",
	UserErrPut:     "用户修改失败",
	UserErrGet:     "用户查询失败",
	UserErrGetList: "用户列表查询失败Config",
}

// 角色类业务状态码 40400 开始
const (
	RoleErrPost int = iota + 40400
	RoleErrDelete
	RoleErrPut
	RoleErrGet
	RoleErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var RoleErrMsg = map[int]string{
	RoleErrPost:    "用户添加失败",
	RoleErrDelete:  "用户删除失败",
	RoleErrPut:     "用户修改失败",
	RoleErrGet:     "用户查询失败",
	RoleErrGetList: "用户列表查询失败Config",
}

// 权限类业务状态码 40400 开始
const (
	PermissionErrPost int = iota + 40400
	PermissionErrDelete
	PermissionErrPut
	PermissionErrGet
	PermissionErrGetList
	// ... 可以继续添加其他集群相关的状态码
)

var PermissionErrMsg = map[int]string{
	PermissionErrPost:    "权限添加失败",
	PermissionErrDelete:  "权限删除失败",
	PermissionErrPut:     "权限修改失败",
	PermissionErrGet:     "权限查询失败",
	PermissionErrGetList: "权限列表查询失败Config",
}

// server服务器内部状态码 50100 开始
const (
	ServerSuccess int = iota + 50100
	ServerNamespaceFailed
	ServerNamespaceFound
	ServerNamespaceError
	// ... 可以继续添加其他集群相关的状态码
)
