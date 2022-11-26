package models

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   int
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
	Role     Role     `gorm:"foreignKey:RoleId;references:Id"`
	Access   []Access `gorm:"many2many:role_access  "` //
}

func (Manager) TableName() string {
	return "manager"
}
