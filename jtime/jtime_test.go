package jtime

import (
	"fmt"
	"navi/common/utils/json"
	"testing"
	"time"
)

type JT struct {
	Tm JsonTime `json:"tm"` //注意大写才能导出
}

func TestJsonTime_UnmarshalJSON(b *testing.T) {

	var t JT = JT{JsonTime{time.Now()}}
	fmt.Printf("系统默认的时间格式: %s \n", t.Tm)

	bt, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("转成常用时间格式: %s \n", string(bt))
	err = json.Unmarshal(bt, &t)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("转换成系统默认格式: %s \n", t.Tm)
}
