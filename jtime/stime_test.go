package jtime

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

type t struct {
	Tm TstampTime
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
	fmt.Printf("转成常用时间格式: %s \n", string(bt))
	err = json.Unmarshal(bt, t2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("转换成时间戳格式: %d \n", t2.Tm)
}
