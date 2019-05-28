package jtime

import "time"

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
