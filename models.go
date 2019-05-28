package common

import (
	"github.com/Gre-Z/common/jtime"
	"time"
)

type Model struct {
	Id        int64 `gorm:"primary_key"`
	CreatedAt jtime.JsonTime
	UpdatedAt jtime.JsonTime
	DeletedAt *time.Time `sql:"index"`
}
