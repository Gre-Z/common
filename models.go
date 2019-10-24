package common

import (
	"github.com/Gre-Z/common/jtime"
	"time"
)

type Model struct {
	Id        int64          `gorm:"primary_key;comment:'主键ID'" json:"id"`
	CreatedAt jtime.JsonTime `gorm:"comment:'创建时间';type:timestamp not null;default:current_timestamp;" json:"created_at"`
	UpdatedAt jtime.JsonTime `gorm:"comment:'更新时间';type:timestamp on update current_timestamp;omitempty;default:current_timestamp;" json:"updated_at"`
	DeletedAt *time.Time     `gorm:"comment:'删除时间'" json:"deleted_at" sql:"index"`
}
