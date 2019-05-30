package common

import (
	"github.com/Gre-Z/common/jtime"
	"time"
)

type Model struct {
	ID        int `gorm:"primary_key" json:"id"`
	CreatedAt jtime.JsonTime
	UpdatedAt jtime.JsonTime
	DeletedAt *time.Time `sql:"index"`
}
