package model

type Role int

const (
	PermissionAdmin       Role = 1 //管理员
	PermissionUser        Role = 2 //普通用户
	PermissionVisitor     Role = 3 //游客
	PermissionDisableUser Role = 4 //被禁用的用户
)

// AuthModel 用户表
type AuthModel struct {
	MODEL
	NickName string `gorm:"size:42" json:"nick_name"`     //昵称
	UserName string `gorm:"size:36" json:"user_name"`     //用户名
	Password string `gorm:"size:128" json:"password"`     //密码
	AvatarID uint   `json:"avatar_id"`                    //头像id
	Email    string `gorm:"size:128" json:"email"`        //邮箱
	Tel      string `gorm:"size:18" json:"tel"`           //手机号
	Addr     string `gorm:"size:64" json:"addr"`          //地址
	Token    string `gorm:"size:64" json:"token"`         //其他平台唯一id
	IP       string `gorm:"size:20" json:"IP"`            //ip地址
	Role     Role   `gorm:"size:4;default:1" json:"role"` //权限
}
