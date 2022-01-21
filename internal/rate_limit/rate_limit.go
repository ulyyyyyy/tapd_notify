package rate_limit

import (
	"fmt"
	"github.com/ulyyyyyy/tapd_notify/internal/helper/wework_notify"
	"github.com/ulyyyyyy/timeboundmap"
	"time"
)

const (
	maxPostPerMinute = 30
	bucketExp        = 1 * time.Minute
)

var (
	boundMap *timeboundmap.TimeBoundMap
	c        = make(chan *count)
)

// Initialize 初始化boundMap
func Initialize() {
	boundMap = timeboundmap.New(
		bucketExp,
		timeboundmap.WithOnCleaned(func(elapsed time.Duration, cleaning, remaining uint64) {
			fmt.Printf("本次共清理 %10d 个, 耗时: %-16s\t(剩余: %10d 个)\n", cleaning, elapsed, remaining)
		}),
	)

	for {
		select {
		case pushAddr := <-c:
			number := (*pushAddr).Number
			receiver := (*pushAddr).Key
			boundMap.Set(receiver, 1, 1*time.Minute, countKey)
			wework_notify.PushMerge(receiver, number)
		}
	}
}

func countKey(key, value interface{}) {
	c <- newCount(key, value)
}

// Check 推送目标值检查是否可以推送，如果超过了频率则合并推送，true为放行
func Check(key string) (rlt bool) {
	num := 1
	if number, ok := boundMap.Get(key); ok {
		num += number.(int)
	}
	boundMap.Set(key, num, 1*time.Minute, countKey)
	return num <= maxPostPerMinute
}

type count struct {
	Key    string
	Number int
}

func newCount(key interface{}, number interface{}) *count {
	return &count{Key: key.(string), Number: number.(int)}
}
