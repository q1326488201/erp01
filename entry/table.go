package entry

import "gorm.io/gorm"

type AuthUser struct {
	gorm.Model
	Username string `json:"username" gorm:"column:username;type:varchar(30);unique;not null"`
	Password string `json:"password" gorm:"column:password;type:varchar(30);not null"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(30);not null"`
	FType    int    `json:"FType" gorm:"column:ftype;type:int;default:0"`
	FNote    string `json:"fnote" gorm:"column:fnote;type:varchar(255)"`
}

func (receiver AuthUser) TableName() string {
	return "auth_user"
}

type AuthRole struct {
	gorm.Model
	FName string `json:"fname" gorm:"column:fname;type:varchar(30);unique;not null"`
	FNote string `json:"fnote" gorm:"column:fnote;type:varchar(255);"`
}

func (receiver AuthRole) TableName() string {
	return "auth_role"
}

type AuthUserRole struct {
	gorm.Model
	UserID int `json:"userID" gorm:"column:userID;type:int;not null"`
	RoleID int `json:"roleID" gorm:"column:roleID;type:int;not null"`
}

func (receiver AuthUserRole) TableName() string {
	return "auth_user_role"
}
