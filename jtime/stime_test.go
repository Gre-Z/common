package jtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type t struct {
	//Tm time.Time
	Tm TstampTime //将原本的 time.Time 替换称TstampTime
}

func TestTstampTime_MarshalJSON(b *testing.T) {
	t2 := &t{
		Tm: TstampTime(time.Now().Unix()),
	}
	bt, err := json.Marshal(t2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("时间戳转成常用时间格式: %s \n", string(bt))
	err = json.Unmarshal(bt, t2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("常用时间格式转换成时间戳格式: %d \n", t2.Tm)
}
