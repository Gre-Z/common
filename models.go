package common

import (
	"github.com/Gre-Z/common/jtime"
	"time"
)

type Model struct {
	Id        uint `gorm:"primary_key" json:"id"`
	CreatedAt jtime.JsonTime
	UpdatedAt jtime.JsonTime
	DeletedAt *time.Time `sql:"index"`
}
