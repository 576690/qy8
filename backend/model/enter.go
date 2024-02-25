package model

import "time"

type MODEL struct {
	ID        uint      `gorm:"primarykey" json:"id"` //主键id
	CreatedAt time.Time `json:"created_at"`           //创建时间
	UpdateAt  time.Time `json:"update_at"`            //更新时间
}
