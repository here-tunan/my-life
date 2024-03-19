package user

type User struct {
	// Id
	Id int64 `json:"id"`
	// 账户
	Account string `json:"account"`
	// 密码
	Password string `json:"password"`
	// 名称
	Name string `json:"name"`
	// 描述
	Desc string `json:"desc"`
	// 头像
	Avatar string `json:"avatar"`
	// 额外信息
	Extra string `json:"extra"`
}
