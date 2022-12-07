package models

type NameUrl struct {
	Username string
	Url      string
}

/* 	db.Table("manager", "access").Select("manager.username, access.url").
Joins("left join role_access on manager.role_id = role_access.manager_id").
Joins("left join access on access.id = role_access.access_id").
Scan(&nameurls) */
//nameurls := []models.NameUrl{}
//models.DB.Raw("SELECT manager.username, access.url FROM manager left join role_access on manager.role_id = role_access.manager_id left join access on access.id = role_access.access_id").Scan(&nameurls)
