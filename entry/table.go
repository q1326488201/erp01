package entry

import "gorm.io/gorm"

type AuthUser struct {
	gorm.Model
	Username  string `json:"username" gorm:"column:username;type:varchar(30);unique;not null"`
	Password  string `json:"password" gorm:"column:password;type:varchar(30);not null"`
	AvatarUrl string `json:"avatar_url" gorm:"column:avatar_url;type:text"`
	Nickname  string `json:"nickname" gorm:"column:nickname;type:varchar(30);not null"`
	FType     int    `json:"FType" gorm:"column:ftype;type:int;default:0"`
	FNote     string `json:"fnote" gorm:"column:fnote;type:varchar(255)"`
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

type AuthMenu struct {
	gorm.Model
	FName         string `json:"name"  gorm:"column:roleID;type:varchar(12)"`               // 菜单名称
	FParentId     int64  `json:"parentId"  gorm:"column:parentId;type:int;"`                // 父菜单ID，一级菜单为0
	FUrl          string `json:"url"  gorm:"column:url;type:varchar(128)"`                  // 菜单URL,类型：1.普通页面（如用户管理， /sysmodel/user） 2.嵌套完整外部页面，以http(s)开头的链接 3.嵌套服务器页面，使用iframe:前缀+目标URL(如SQL监控， iframe:/druid/login.html, iframe:前缀会替换成服务器地址)
	FPerms        string `json:"perms" gorm:"column:perms;type:varchar(128)"`               // 授权(多个用逗号分隔，如：sysmodel:user:add,sysmodel:user:edit)
	FIcon         string `json:"icon"  gorm:"column:icon;type:varchar(128)"`                // 菜单图标
	FOrderNum     int64  `json:"orderNum"  gorm:"column:orderNum;type:int"`                 // 排序
	FVuePath      string `json:"vuePath"  gorm:"column:vuePath;type:varchar(20)"`           // vue系统的path
	FVueComponent string `json:"vueComponent"  gorm:"column:vueComponent;type:varchar(20)"` // vue的页面
	FVueRedirect  string `json:"vueRedirect"  gorm:"column:vueRedirect;type:varchar(20)"`   // vue的路由重定向
}

func (receiver AuthMenu) TableName() string {
	return "auth_menu"
}
