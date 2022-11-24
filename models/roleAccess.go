package models

// 一个一对多的表
// 用于表示access和Role之间的连接
type RoleAccess struct {
	AccessId int
	RoleId   int
}

func (RoleAccess) TableName() string {
	return "role_access"
}
