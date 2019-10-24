//redis 限流功能
package redis

import (
	"errors"
	"time"
)

func (this *RD) Limit(key string, totalTime time.Duration, count int64, unitTime time.Duration) (num int64, err error) {
	if count <= 0 {
		return 0, errors.New("Count must be greater than 0.")
	}

	rd := this.RedisNew()
	// 判断至少隔多久执行一次更新操作 如果是0则不做限制
	if unitTime != 0 {
		cmd := rd.Get(key + "_timer")
		if cmd.Err() == nil {
			return 0, errors.New("Number of deliverables exceeding unit time.")
		} else {
			rd.Set(key+"_timer", 0, unitTime)
		}
	}
	// 限制这个时间区间内只允许操作的次数
	get := rd.Get(key)
	if get.Err() != nil {
		rd.Set(key, 0, totalTime)
	} else {
		l, err := get.Uint64()
		if err != nil {
			return 0, err
		}
		num = int64(l)
		if !(0 <= num && num < count) {
			return num, errors.New("exceeding the limit")
		}
	}
	val := rd.Incr(key).Val()
	return val, nil
}
