package member

import "gorm.io/gorm"

type Config struct {
	Auth       bool
	Permission bool
	DB         *gorm.DB
}

type Member struct {
	Auth          bool
	Permission    bool
	AuthFlg       bool
	PermissionFlg bool
	DBString      *gorm.DB
}
