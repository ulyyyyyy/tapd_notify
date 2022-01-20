package rate_limit

import (
	"fmt"
	"github.com/ulyyyyyy/timeboundmap"
	"time"
)

const (
	maxPostPerMinute = 30
	bucketExp        = 1 * time.Minute
)

var (
	unsentContentMap = make(map[string]int) // 未发送的消息
	boundMap         *timeboundmap.TimeBoundMap
	c                = make(chan string)
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
			fmt.Println(pushAddr)
		}
	}
}

// Check 推送目标值检查是否可以推送，如果超过了频率则合并推送
func Check(key string) (rlt bool) {
	c := make(chan string)

	num := 1
	if number, ok := boundMap.Get(key); ok {
		num += number.(int)
	}
	boundMap.Set(key, num, 1*time.Second, func(key, value interface{}) {
		fmt.Println(key, value)
		c <- key.(string)
		delete(unsentContentMap, key.(string))
	})
	rlt = num <= maxPostPerMinute
	if !rlt {

	}
	return
}

// GetUnsentLen 检查合并推送数据的条数
func GetUnsentLen(key string) (len int) {
	return unsentContentMap[key]
}
