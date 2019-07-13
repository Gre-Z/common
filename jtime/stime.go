package jtime

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type TstampTime int64

func (t *TstampTime) UnmarshalJSON(data []byte) (err error) {
	time.ParseInLocation(`"`+TimeFormat+`"`, string(data), time.Local)
	return
}

func (t TstampTime) MarshalJSON() ([]byte, error) {
	tint := int64(t)
	if tint == 0 {
		return []byte(`""`), nil
	}
	tf := time.Unix(tint, 0).Format(`"` + TimeFormat + `"`)
	return []byte(tf), nil
}

func (t TstampTime) Value() (driver.Value, error) {
	var tm int64
	if t == 0 {
		tm = time.Now().Unix()
	} else {
		tm = int64(t)
	}
	return tm, nil
}

func (t *TstampTime) Scan(v interface{}) error {
	value, ok := v.(int64)
	if ok {
		*t = TstampTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
