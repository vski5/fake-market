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
	Role     Role   `gorm:"foreignKey:RoleId;references:Id"`
	Access   Access `gorm:"many2many:role_access;references:Id;foreignKey:RoleId"` //;references:AccessId;foreignKey:RoleId
}

func (Manager) TableName() string {
	return "manager"
}
