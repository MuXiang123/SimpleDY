// Package dao
// @author    : MuXiang123
// @time      : 2022/6/11 16:31
package dao

// Follow 关注数据库映射
// HostId 关注者
// GuestId 被关注者
type Follow struct {
	Id      uint64
	HostId  uint64
	GuestId uint64
}

// TableName 设置Follow结构体对应数据库表名。
func (Follow) TableName() string {
	return "follow"
}
