package dao

// User 对应数据库User表结构的结构体
type User struct {
	Id            uint64 //自增主键
	Name          string //昵称
	Username      string //用户名
	Password      string //密码
	followCount   uint64 //关注数
	followerCount uint64 //粉丝数
}

// TableName 修改表名映射
func (user User) TableName() string {
	return "user"
}

//用户注册参数
type UserRegisterParam struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
	Name     string `form:"name" json:"name"`
}

//注册返回信息
//type UserRegisterResponse struct {
//	Response
//	UserId uint64 `json:"id"`
//}

//用户登录参数
type UserLoginParam struct {
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}
